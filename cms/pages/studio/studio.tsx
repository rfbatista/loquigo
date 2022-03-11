import React from 'react';
import {
  Container,
  Content,
  Header,
  Sidebar,
  Loader,
  toaster,
  Message,
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
      .catch((error) => toaster.push(message, { placement: 'topCenter' }));
  };
  const message = (
    <Message showIcon type='error' header='Error'>
      Não foi posivel salvar.
    </Message>
  );

  return (
    <>
      <Container className='my-5'>
        <Panel updateBot={send} isLoading={isUpdating} />
      </Container>
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
        className='place-items-center w-full'
        style={{ height: '100vh', maxWidth: '1000px', minWidth: '400px' }}
      >
        <YamlEditor data={data} />
      </Container>
    </div>
  );
};

export default Studio;
