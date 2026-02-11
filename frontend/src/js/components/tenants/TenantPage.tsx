// Copyright 2024 Northern.tech AS
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
import { useCallback, useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { Add as AddIcon } from '@mui/icons-material';
import { Button, Typography } from '@mui/material';
import { makeStyles } from 'tss-react/mui';

import { getSpLimits, getTenantsList } from '@northern.tech/store/selectors';
import { AppDispatch } from '@northern.tech/store/store';
import { getTenants } from '@northern.tech/store/thunks';
import { toggle } from '@northern.tech/utils/helpers';

import { DeviceLimit } from '../header/DeviceNotifications';
import { TenantCreateForm } from './TenantCreateForm';
import { TenantList } from './TenantList';

interface TenantsEmptyStateProps {
  openModal: () => void;
}
const TenantsEmptyState = (props: TenantsEmptyStateProps) => {
  const { openModal } = props;
  const dispatch = useDispatch<AppDispatch>();
  useEffect(() => {
    dispatch(getTenants());
  }, [dispatch]);
  return (
    <div className="dashboard-placeholder">
      <p>You are not currently managing any tenants. </p>
      <p>
        <a onClick={openModal}>Add a tenant</a> to get started.
      </p>
    </div>
  );
};

const useStyles = makeStyles()(theme => ({
  limit: {
    border: `1px solid #E0E0E0`,
    maxWidth: '726px'
  }
}));

export const TenantPage = () => {
  const [showCreate, setShowCreate] = useState<boolean>(false);

  const { classes } = useStyles();

  const { tenants } = useSelector(getTenantsList);
  const spLimits = useSelector(getSpLimits);

  const onToggleCreation = useCallback(() => setShowCreate(toggle), []);
  return (
    <div className="padding-right">
      <Typography variant="h5">Tenant management</Typography>
      <Typography variant="subtitle1" className="margin-top-small">
        Device limits
      </Typography>
      <div className={`full-width flexbox`}>
        {Object.values(spLimits)
          .map(limit => (
            <div className={`full-width padding-small margin-right-small ${classes.limit}`}>
              <DeviceLimit serviceProvider total={limit.current} limit={limit.limit} type={limit.name} />
            </div>
          ))}
      </div>
      <div className="flexbox full-width space-between">
        <Typography variant="subtitle1" className="margin-top-small">
          Tenants
        </Typography>
        <Button className="margin-top-small" variant="contained" onClick={onToggleCreation}>
          Create a tenant
        </Button>
      </div>

      {tenants.length ? <TenantList /> : <TenantsEmptyState openModal={onToggleCreation} />}
      {showCreate && <TenantCreateForm open={showCreate} onCloseClick={onToggleCreation} />}
    </div>
  );
};
