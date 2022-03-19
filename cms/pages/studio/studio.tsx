import React from 'react';
import {
  Container,
  Content,
  Header,
  Sidebar,
  Loader,
  toaster,
  Message,
  Navbar,
  Nav,
  Dropdown,
	InputPicker,
  FlexboxGrid,
} from 'rsuite';
import dynamic from 'next/dynamic';
import Panel from './panel';
import { useGetBotQuery, useUpdateBotMutation } from 'services/loquiapi';
const MonacoEditor = dynamic(import('react-monaco-editor'), { ssr: false });

const YamlEditor = ({ data }) => {
  const [postBody, setPostBody] = React.useState(data);
  const [updateBot, { isLoading: isUpdating }] = useUpdateBotMutation();
  const send = () => {
    updateBot(postBody)
      .unwrap()
      .then(() => toaster.push(sucessMessage, { placement: 'topCenter' }))
      .catch((error) => toaster.push(errorMessage, { placement: 'topCenter' }));
  };
  const errorMessage = (
    <Message showIcon type='error' header='Erro'>
      Não foi posivel salvar.
    </Message>
  );
  const sucessMessage = (
    <Message showIcon type='success' header='Sucesso'>
      Atualizava salva com sucesso.
    </Message>
  );
  return (
    <>
      <Header className='w-full'>
        <Navbar appearance='inverse'>
          <Nav>
            <Nav.Item>
              <Panel updateBot={send} isLoading={isUpdating} />
            </Nav.Item>
            <Nav.Item>
              <InputPicker style={{ width: 224 }} />
            </Nav.Item>
          </Nav>
        </Navbar>
      </Header>
      <Content className='w-full h-full'>
        <FlexboxGrid className='h-full'>
          <FlexboxGrid.Item className='h-full' colspan={12}>
            <Container
              className={'h-full w-full'}
              style={{ maxWidth: '1000px', minWidth: '400px' }}
            >
              <MonacoEditor
                editorDidMount={() => {
                  /* @ts-ignore */
                  window.MonacoEnvironment.getWorkerUrl = (
                    _moduleId: string,
                    label: string
                  ) => {
                    if (label === 'json') return '_next/static/json.worker.js';
                    if (label === 'css') return '_next/static/css.worker.js';
                    if (label === 'html') return '_next/static/html.worker.js';
                    if (label === 'yaml') return '_next/static/yaml.worker.js';
                    if (label === 'typescript' || label === 'javascript')
                      return '_next/static/ts.worker.js';
                    return '_next/static/editor.worker.js';
                  };
                }}
                language='yaml'
                theme='vs-dark'
                value={postBody}
                options={{
                  minimap: {
                    enabled: false,
                  },
                }}
                onChange={setPostBody}
              />
            </Container>
          </FlexboxGrid.Item>
          <FlexboxGrid.Item colspan={12}></FlexboxGrid.Item>
        </FlexboxGrid>
      </Content>
    </>
  );
};

const Studio = () => {
  const { data, isFetching, isLoading } = useGetBotQuery('teste');
  if (isFetching || isLoading)
    return <Loader backdrop inverse center content='loading...' vertical />;
  return (
    <div className='grid place-items-center h-screen'>
      <Container
        className='place-items-center w-full h-full'
        style={{ height: '100vh' }}
      >
        <YamlEditor data={data} />
      </Container>
    </div>
  );
};

export default Studio;
