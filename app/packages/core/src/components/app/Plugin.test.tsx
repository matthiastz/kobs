import { render as _render, RenderResult, screen } from '@testing-library/react';
import { FunctionComponent } from 'react';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

import Plugin from './Plugin';

import { IPluginContext, IPluginInstance, PluginContext } from '../../context/PluginContext';

describe('Plugin', () => {
  const render = (path: string, pluginContext: Partial<IPluginContext>): RenderResult => {
    const defaultPluginContext = {
      getClusters: () => [],
      getInstance: (id: string) => undefined,
      getPlugin: (type: string) => undefined,
      getPluginTypes: () => [],
      instances: [],
    };
    return _render(
      <MemoryRouter initialEntries={[path]}>
        <PluginContext.Provider value={{ ...defaultPluginContext, ...pluginContext }}>
          <Routes>
            <Route path="/plugins/:cluster/:type/:name" element={<Plugin />} />
          </Routes>
        </PluginContext.Provider>
      </MemoryRouter>,
    );
  };

  it('should render the plugin page with the given instance props', async () => {
    const TestPage: FunctionComponent<{ instance: IPluginInstance }> = ({ instance: { cluster, id, name, type } }) => {
      return (
        <>
          my plugin with props cluster={cluster} id={id} name={name} type={type}
        </>
      );
    };

    const instance = {
      cluster: 'dev',
      id: '/cluster/dev/bar/bar-instance',
      name: 'bar-instance',
      type: 'bar',
    };

    render('/plugins/dev/bar/bar-instance', {
      getInstance: () => instance,
      getPlugin: () => ({
        page: TestPage,
        type: 'bar',
      }),
    });

    expect(
      screen.getByText(
        `my plugin with props cluster=${instance.cluster} id=${instance.id} name=${instance.name} type=${instance.type}`,
      ),
    ).toBeInTheDocument();
  });
});