import { useSelector } from 'react-redux';
import {
  Header,
  Navbar,
  Input,
  Nav,
  PanelGroup,
  Panel,
  InputPicker,
  Message,
  toaster,
  Button,
} from 'rsuite';
import {
  useGetBotVersionsQuery,
  useGetBotListQuery,
  useUpdateBotMutation,
} from 'services/loquiapi';
import { getActiveBotYaml } from 'store/flow';
import { useState } from 'react';
import SystemModal from 'components/modal';

const formatVersionToInputPicker = (result) => {
  return result?.data
    ? result.data.map((version) => {
        return { label: 'Eugenia', value: 'Eugenia', role: 'Master' };
      })
    : [];
};

const formatBotsToInputPicker = (result) => {
  return result?.data
    ? result.data.map((bot) => {
        return { label: bot.name, value: bot.id, role: 'Master' };
      })
    : [];
};

const BotPicker = () => {
  const { data: botList } = useGetBotListQuery('');
  return (
    <>
      <div className='mb-2'> Lista de bots criados </div>
      <InputPicker
        placeholder='Selecionar'
        size='xs'
        data={formatBotsToInputPicker(botList)}
				onChange={((data)=>console.log(data))}
      />
    </>
  );
};

const VersionPicker = (botId) => {
  const { data: botVersions, isFetching, isLoading } = useGetBotVersionsQuery(
    botId
  );
  return (
    <>
      <label className='m-2'> Criados </label>
      <InputPicker size='xs' data={formatVersionToInputPicker(botVersions)} />
    </>
  );
};

const EditorPanel = () => {
  const [openCreateBotModal, setCreateBotModal] = useState(false);
  const botYaml = useSelector(getActiveBotYaml);
  const [updateBot, { isLoading: isUpdating }] = useUpdateBotMutation();
  const send = () => {
    updateBot(botYaml)
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
      <SystemModal
        title='Criar bot'
        open={openCreateBotModal}
        handleClose={() => setCreateBotModal(false)}
      >
        <Input className='my-2' placeholder='Default Input' />
      </SystemModal>
      <PanelGroup className='h-fit'>
        <Panel header='Bots' defaultExpanded>
          <BotPicker />
          <Button
            appearance='primary'
            className='bg-blue-600 text-white my-3'
            size='sm'
            block
            onClick={() => setCreateBotModal(true)}
          >
            Criar novo bot
          </Button>
        </Panel>
        <Panel header='Publicação'>
          <Button
            className={
							botYaml ? 
							'bg-blue-600 text-white my-3' : 
							'bg-gray-500 text-white my-3'
						}
            size='sm'
            block
            onClick={() => setCreateBotModal(true)}
            disabled={botYaml ? false : true}
          >
            Publicar bot
          </Button>
        </Panel>
        <Panel header='Versões'></Panel>
      </PanelGroup>
    </>
  );
};

export default EditorPanel;
