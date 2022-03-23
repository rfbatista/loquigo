import EditorPanel from 'components/panel';
import YamlEditor from 'components/yaml-editor';
import type { NextPage } from 'next';
import React from 'react';
import { Container, Loader, Sidebar } from 'rsuite';
import { useGetBotYamlQuery } from 'services/loquiapi';
import SidePanel from '../components/side-panel/index';

const Home: NextPage = () => {
  const botId = 'teste';
  const { data, isFetching, isLoading } = useGetBotYamlQuery(botId);
  if (isFetching || isLoading)
    return <Loader backdrop inverse center content='loading...' vertical />;
  return (
    <div className='grid place-items-center h-screen'>
      <Container
        className='place-items-center w-full h-full'
        style={{ height: '100vh' }}
      >
        <Sidebar>
          <SidePanel />
        </Sidebar>
        <Container>
          <YamlEditor data={data} botId={botId} />
        </Container>
      </Container>
    </div>
  );
};

export default Home;
