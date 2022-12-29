import React, { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

import { IPluginInstance, ITimes, PageContentSection, PageHeaderSection, PluginPageTitle } from '@kobsio/shared';
import { IOptions } from '../../utils/interfaces';
import Logs from './Logs';
import LogsPageActions from './LogsPageActions';
import LogsToolbar from './LogsToolbar';
import { defaultDescription } from '../../utils/constants';
import { getInitialOptions } from '../../utils/helpers';

interface ILogsPageProps {
  instance: IPluginInstance;
}

const LogsPage: React.FunctionComponent<ILogsPageProps> = ({ instance }: ILogsPageProps) => {
  const location = useLocation();
  const navigate = useNavigate();
  const [options, setOptions] = useState<IOptions>();

  // changeOptions is used to change the options. Besides setting a new value for the options state we also reflect the
  // options in the current url.
  const changeOptions = (opts: IOptions): void => {
    const fields = opts.fields ? opts.fields.map((field) => `&field=${field}`) : [];

    navigate(
      `${location.pathname}?query=${encodeURIComponent(opts.query)}&order=${opts.order}&orderBy=${opts.orderBy}&time=${
        opts.times.time
      }&timeEnd=${opts.times.timeEnd}&timeStart=${opts.times.timeStart}${fields.length > 0 ? fields.join('') : ''}`,
    );
  };

  // selectField is used to add a field as parameter, when it isn't present and to remove a fields from as parameter,
  // when it is already present via the changeOptions function.
  const selectField = (field: string): void => {
    if (options) {
      let tmpFields: string[] = [];
      if (options.fields) {
        tmpFields = [...options.fields];
      }

      if (tmpFields.includes(field)) {
        tmpFields = tmpFields.filter((f) => f !== field);
      } else {
        tmpFields.push(field);
      }

      changeOptions({ ...options, fields: tmpFields });
    }
  };

  const changeFieldOrder = (oldIndex: number, newIndex: number): void => {
    if (options && options.fields) {
      const tmpFields = [...options.fields];
      const tmpField = tmpFields[oldIndex];

      tmpFields[oldIndex] = tmpFields[newIndex];
      tmpFields[newIndex] = tmpField;

      changeOptions({ ...options, fields: tmpFields });
    }
  };

  // addFilter adds the given filter as string to the query, so that it can be used to filter down an existing query.
  const addFilter = (filter: string): void => {
    if (options) {
      changeOptions({ ...options, query: `${options.query} ${filter}` });
    }
  };

  // changeTime changes the selected time range. This can be used to change the time outside of the toolbar, e.g. by
  // selecting a time range in the chart.
  const changeTime = (times: ITimes): void => {
    if (options) {
      changeOptions({ ...options, times: times });
    }
  };

  // changeOrder changes the order parameters for a query.
  const changeOrder = (order: string, orderBy: string): void => {
    if (options) {
      changeOptions({ ...options, order: order, orderBy: orderBy });
    }
  };

  useEffect(() => {
    setOptions((prevOptions) => getInitialOptions(location.search, !prevOptions));
  }, [location.search]);

  if (!options) {
    return null;
  }

  return (
    <React.Fragment>
      <PageHeaderSection
        component={
          <PluginPageTitle
            satellite={instance.satellite}
            name={instance.name}
            description={instance.description || defaultDescription}
            actions={<LogsPageActions instance={instance} query={options.query || ''} />}
          />
        }
      />

      <PageContentSection
        hasPadding={true}
        hasDivider={true}
        toolbarContent={<LogsToolbar options={options} setOptions={changeOptions} />}
        panelContent={undefined}
      >
        {options.query.length > 0 ? (
          <Logs
            instance={instance}
            fields={options.fields}
            query={options.query}
            order={options.order}
            orderBy={options.orderBy}
            addFilter={addFilter}
            changeTime={changeTime}
            changeOrder={changeOrder}
            selectField={selectField}
            changeFieldOrder={changeFieldOrder}
            times={options.times}
          />
        ) : (
          <div></div>
        )}
      </PageContentSection>
    </React.Fragment>
  );
};

export default LogsPage;
