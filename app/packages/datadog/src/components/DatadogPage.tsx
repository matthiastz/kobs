import { IPluginPageProps } from '@kobsio/core';
import { FunctionComponent } from 'react';
import { Route, Routes } from 'react-router-dom';

import LogsPage from './LogsPage';

const KLogsPage: FunctionComponent<IPluginPageProps> = (props) => {
  return (
    <Routes>
      <Route path="/" element={<LogsPage {...props} />} />
    </Routes>
  );
};

export default KLogsPage;
