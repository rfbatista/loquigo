import React from 'react';
import {
  toaster,
  Message,
  Header,
  Navbar,
  Nav,
  InputPicker,
  Content,
  FlexboxGrid,
  Container,
} from 'rsuite';
import {
  useGetBotVersionsQuery,
  useUpdateBotMutation,
} from 'services/loquiapi';
import Panel from '../panel/panel';
import dynamic from 'next/dynamic';
const MonacoEditor = dynamic(import('react-monaco-editor'), { ssr: false });

const formatVersionToInputPicker = (data) => {
  return data
    ? data.map((version) => {
        return { label: 'Eugenia', value: 'Eugenia', role: 'Master' };
      })
    : [];
};

const formatBotsToInputPicker = (data) => {
  return data
    ? data.map((version) => {
        return { label: 'Eugenia', value: 'Eugenia', role: 'Master' };
      })
    : [];
};

type YamlEditorProps = {
  data: string;
  botId: string;
};
const YamlEditor: React.FC<YamlEditorProps> = ({ data, botId }) => {
  const [postBody, setPostBody] = React.useState(data);
  const { data: botVersions, isFetching, isLoading } = useGetBotVersionsQuery(
    botId
  );
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
              <InputPicker
                data={formatVersionToInputPicker(botVersions)}
                style={{ width: 224 }}
              />
            </Nav.Item>
            <Nav.Item>
              <InputPicker
                data={formatVersionToInputPicker(botVersions)}
                style={{ width: 224 }}
              />
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

export default YamlEditor;
