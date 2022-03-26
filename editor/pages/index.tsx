import YamlEditor from 'components/yaml-editor';
import type { NextPage } from 'next';
import React from 'react';

import { Container, Loader, Sidebar } from 'rsuite';

import SidePanel from '../components/side-panel/index';

const Home: NextPage = () => {


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
          <YamlEditor />
        </Container>
      </Container>
    </div>
  );
};

export default Home;
