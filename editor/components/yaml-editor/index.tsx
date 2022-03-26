import React from 'react';
import { Content, FlexboxGrid, Container, Loader } from 'rsuite';
import dynamic from 'next/dynamic';
import { useSelector } from 'react-redux';
import { useGetBotYamlQuery } from 'services/loquiapi';
import { selectActiveBotId } from 'store/bot';
const MonacoEditor = dynamic(import('react-monaco-editor'), { ssr: false });

type YamlEditorProps = {};

const YamlEditor: React.FC<YamlEditorProps> = () => {
  const botId = useSelector(selectActiveBotId);
  const { data, isFetching, isLoading } = useGetBotYamlQuery(botId);
  return (
    <>
      {isFetching || isLoading ? (
        <Loader backdrop inverse center content='loading...' vertical />
      ) : (
        <EditorBody data={data} />
      )}
    </>
  );
};

const EditorBody = ({ data }) => {
  const [postBody, setPostBody] = React.useState(data);
  return (
    <>
      <Content
        className='w-full h-full'
        style={{ height: '100vh', width: '100%' }}
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
      </Content>
    </>
  );
};

export default YamlEditor;
