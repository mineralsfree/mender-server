// Copyright 2015 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
import React, { useCallback, useEffect, useState } from 'react';
import { useIdleTimer, workerTimers } from 'react-idle-timer';
import { Provider, useDispatch, useSelector } from 'react-redux';
import { BrowserRouter, useLocation, useNavigate } from 'react-router-dom';

import createCache from '@emotion/cache';
import { CacheProvider } from '@emotion/react';
import { CssBaseline, GlobalStyles, ThemeProvider, createTheme, styled } from '@mui/material';
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { makeStyles } from 'tss-react/mui';

import SharedSnackbar from '@northern.tech/common-ui/SharedSnackbar';
import ConfirmDismissHelptips from '@northern.tech/common-ui/dialogs/ConfirmDismissHelpTips';
import DeviceConnectionDialog from '@northern.tech/common-ui/dialogs/DeviceConnectionDialog';
import FeedbackDialog from '@northern.tech/common-ui/dialogs/Feedback';
import StartupNotificationDialog from '@northern.tech/common-ui/dialogs/StartupNotification';
import storeActions from '@northern.tech/store/actions';
import { getSessionInfo, maxSessionAge, updateMaxAge } from '@northern.tech/store/auth';
import { TIMEOUTS } from '@northern.tech/store/constants';
import { getCurrentSession, getCurrentUser, getIsDarkMode, getIsServiceProvider, getSnackbar, getTrackerCode } from '@northern.tech/store/selectors';
import { store } from '@northern.tech/store/store';
import { parseEnvironmentInfo } from '@northern.tech/store/storehooks';
import { logoutUser } from '@northern.tech/store/thunks';
import { toggle } from '@northern.tech/utils/helpers';
import Cookies from 'universal-cookie';

import ErrorBoundary from '../ErrorBoundary';
import { PrivateRoutes, PrivateSPRoutes, PublicRoutes } from '../config/routes';
import { dark as darkTheme, light as lightTheme } from '../themes/Mender';
import Tracking from '../tracking';
import Footer from './Footer';
import LeftNav from './LeftNav';
import SearchResult from './SearchResult';
import Uploads from './Uploads';
import Header from './header/Header';

const { receivedActivationCode, setShowConnectingDialog, setSnackbar } = storeActions;

const cache = createCache({ key: 'mui', prepend: true });

const activationPath = '/activate';
const trackingBlacklist = [/\/password\/.+/i];
const timeout = maxSessionAge * 1000; // 15 minutes idle time
const cookies = new Cookies();

const reducePalette =
  prefix =>
  (accu, [key, value]) => {
    if (value instanceof Object) {
      return {
        ...accu,
        ...Object.entries(value).reduce(reducePalette(`${prefix}-${key}`), {})
      };
    } else {
      accu[`${prefix}-${key}`] = value;
    }
    return accu;
  };

const cssVariables = ({ theme: { palette } }) => {
  const muiVariables = Object.entries(palette).reduce(reducePalette('--mui'), {});
  return {
    '@global': {
      ':root': {
        ...muiVariables,
        '--mui-overlay': palette.grey[400]
      }
    }
  };
};

const WrappedBaseline = styled(CssBaseline)(cssVariables);

const useStyles = makeStyles()(() => ({
  public: {
    display: 'grid',
    gridTemplateRows: 'max-content 1fr max-content',
    height: '100vh',
    '.content': {
      alignSelf: 'center',
      justifySelf: 'center'
    }
  }
}));

export const AppRoot = () => {
  const [showSearchResult, setShowSearchResult] = useState(false);
  const navigate = useNavigate();
  const { pathname = '', hash } = useLocation();

  const dispatch = useDispatch();
  const { id: currentUser } = useSelector(getCurrentUser);
  const showDismissHelptipsDialog = useSelector(state => !state.onboarding.complete && state.onboarding.showTipsDialog);
  const showDeviceConnectionDialog = useSelector(state => state.users.showConnectDeviceDialog);
  const showStartupNotification = useSelector(state => state.users.showStartupNotification);
  const showFeedbackDialog = useSelector(state => state.users.showFeedbackDialog);
  const snackbar = useSelector(getSnackbar);
  const trackingCode = useSelector(getTrackerCode);
  const isDarkMode = useSelector(getIsDarkMode);
  const { token: storedToken } = getSessionInfo();
  const { expiresAt, token = storedToken } = useSelector(getCurrentSession);

  const trackLocationChange = useCallback(
    pathname => {
      let page = pathname;
      // if we're on page whose path might contain sensitive device/ group/ deployment names etc. we sanitize the sent information before submission
      if (page.includes('=') && (page.startsWith('/devices') || page.startsWith('/deployments'))) {
        const splitter = page.lastIndexOf('/');
        const filters = page.slice(splitter + 1);
        const keyOnlyFilters = filters.split('&').reduce((accu, item) => `${accu}:${item.split('=')[0]}&`, ''); // assume the keys to filter by are not as revealing as the values things are filtered by
        page = `${page.substring(0, splitter)}?${keyOnlyFilters.substring(0, keyOnlyFilters.length - 1)}`; // cut off the last & of the reduced filters string
      } else if (page.startsWith(activationPath)) {
        dispatch(receivedActivationCode(page.substring(activationPath.length + 1)));
        navigate('/settings/my-profile', { replace: true });
      } else if (trackingBlacklist.some(item => !!page.match(item))) {
        return;
      }
      Tracking.pageview(page);
    },
    [dispatch, navigate]
  );

  useEffect(() => {
    dispatch(parseEnvironmentInfo());
    if (!trackingCode) {
      return;
    }
    if (!cookies.get('_ga')) {
      Tracking.cookieconsent().then(({ trackingConsentGiven }) => {
        if (trackingConsentGiven) {
          Tracking.initialize(trackingCode);
          Tracking.pageview();
        }
      });
    } else {
      Tracking.initialize(trackingCode);
    }
  }, [dispatch, trackingCode]);

  useEffect(() => {
    if (!(trackingCode && cookies.get('_ga'))) {
      return;
    }
    trackLocationChange(pathname);
  }, [pathname, trackLocationChange, trackingCode]);

  useEffect(() => {
    trackLocationChange(pathname);
    // the following is added to ensure backwards capability for hash containing routes & links (e.g. /ui/#/devices => /ui/devices)
    if (hash) {
      navigate(hash.substring(1));
    }
  }, [hash, navigate, pathname, trackLocationChange]);

  const updateExpiryDate = useCallback(() => updateMaxAge({ expiresAt, token }), [expiresAt, token]);

  const onIdle = useCallback(() => {
    if (!!expiresAt && currentUser) {
      // logout user and warn
      return dispatch(logoutUser())
        .catch(updateExpiryDate)
        .then(() => {
          navigate('//'); // double / to ensure the logged out URL conforms to `/ui/` in order to not trigger a redirect and potentially use state
          // async snackbar setting to ensure the login screen has loaded as the snackbar might be cleared by other actions otherwise
          setTimeout(() => dispatch(setSnackbar('Your session has expired. You have been automatically logged out due to inactivity.')), TIMEOUTS.oneSecond);
        });
    }
  }, [currentUser, dispatch, expiresAt, navigate, updateExpiryDate]);

  useIdleTimer({ crossTab: true, onAction: updateExpiryDate, onActive: updateExpiryDate, onIdle, syncTimers: 400, timeout, timers: workerTimers });

  const onToggleSearchResult = () => setShowSearchResult(toggle);

  const theme = createTheme(isDarkMode ? darkTheme : lightTheme);

  const { classes } = useStyles();
  const globalCssVars = cssVariables({ theme })['@global'];

  const dispatchedSetSnackbar = useCallback(message => dispatch(setSnackbar(message)), [dispatch]);
  const isSP = useSelector(getIsServiceProvider);

  return (
    <ThemeProvider theme={theme}>
      <WrappedBaseline enableColorScheme />
      <GlobalStyles styles={globalCssVars} />
      <>
        {token ? (
          <div id="app">
            <Header isDarkMode={isDarkMode} />
            <LeftNav />
            <div className="rightFluid container">
              <ErrorBoundary>
                <SearchResult onToggleSearchResult={onToggleSearchResult} open={showSearchResult} />
                {isSP ? <PrivateSPRoutes /> : <PrivateRoutes />}
              </ErrorBoundary>
            </div>
            {showDismissHelptipsDialog && <ConfirmDismissHelptips />}
            {showDeviceConnectionDialog && <DeviceConnectionDialog onCancel={() => dispatch(setShowConnectingDialog(false))} />}
            {showStartupNotification && <StartupNotificationDialog />}
            {showFeedbackDialog && <FeedbackDialog />}
          </div>
        ) : (
          <div className={classes.public}>
            <PublicRoutes />
            <Footer />
          </div>
        )}
        <SharedSnackbar snackbar={snackbar} setSnackbar={dispatchedSetSnackbar} />
        <Uploads />
      </>
    </ThemeProvider>
  );
};

export const AppProviders = ({ basename = 'ui' }) => (
  <React.StrictMode>
    <Provider store={store}>
      <CacheProvider value={cache}>
        <LocalizationProvider dateAdapter={AdapterDayjs}>
          <ErrorBoundary>
            <BrowserRouter basename={basename}>
              <AppRoot />
            </BrowserRouter>
          </ErrorBoundary>
        </LocalizationProvider>
      </CacheProvider>
    </Provider>
  </React.StrictMode>
);

export default AppRoot;
