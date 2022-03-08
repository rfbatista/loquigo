import React from 'react';
import { Container, Content, Header, Sidebar } from 'rsuite';
import dynamic from 'next/dynamic';
const MonacoEditor = dynamic(import('react-monaco-editor'), { ssr: false });

const Studio = () => {
  const [postBody, setPostBody] = React.useState('');
	console.log(postBody)
  return (
    <div className='grid place-items-center h-screen'>
      <Container style={{ height: '100vh' }}>
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
              if (label === 'typescript' || label === 'javascript')
                return '_next/static/ts.worker.js';
              return '_next/static/editor.worker.js';
            };
          }}
          width='800'
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
