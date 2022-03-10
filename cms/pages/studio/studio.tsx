import React from 'react';
import { Container, Content, Header, Sidebar } from 'rsuite';
import dynamic from 'next/dynamic';
import Panel from './panel';
import { useUpdateBotMutation } from 'services/loquiapi';
const MonacoEditor = dynamic(import('react-monaco-editor'), { ssr: false });

const Studio = () => {
  const [postBody, setPostBody] = React.useState('');
  const [updateBot] = useUpdateBotMutation();
  const update = () => {
    updateBot(postBody);
  };
  return (
    <div className='grid place-items-center h-screen'>
      <Container className='my-5'>
        <Panel updateBot={update} />
      </Container>
      <Container
        className='place-items-center w-full'
        style={{ height: '100vh', maxWidth: '1000px', minWidth: '400px' }}
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
    </div>
  );
};

export default Studio;
