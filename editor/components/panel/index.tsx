import { useDispatch, useSelector } from 'react-redux';
import {
  Input,
  PanelGroup,
  Panel,
  InputPicker,
  Message,
  toaster,
  Button,
  Loader,
} from 'rsuite';
import {
  useGetBotVersionsQuery,
  useGetBotListQuery,
  useUpdateBotMutation,
  useCreateBotMutation,
} from 'services/loquiapi';
import { getActiveBotYaml } from 'store/flow';
import { useState } from 'react';
import SystemModal from 'components/modal';
import { setActiveBotId } from 'store/bot';
import { useForm } from 'react-hook-form';
import { TextField } from '@mui/material';
import { v4 as uuidv4 } from 'uuid';

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
  const dipatch = useDispatch();
  const { data: botList } = useGetBotListQuery('');
  return (
    <>
      <div className='mb-2'> Lista de bots criados </div>
      <InputPicker
        placeholder='Selecionar'
        size='xs'
        data={formatBotsToInputPicker(botList)}
        onChange={(data) => dipatch(setActiveBotId(data))}
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

const BotModal = ({ openCreateBotModal, setCreateBotModal }) => {
  const [createBot, { isLoading }] = useCreateBotMutation();
  const { register, handleSubmit } = useForm();
  const onSubmit = (data) =>
    createBot({ id: uuidv4(), name: data.botName })
      .unwrap()
      .then(() => toaster.push(sucessMessage, { placement: 'topCenter' }))
      .catch((error) => toaster.push(errorMessage, { placement: 'topCenter' }));

  const errorMessage = (
    <Message showIcon type='error' header='Erro'>
      Não foi posivel criar o bot.
    </Message>
  );
  const sucessMessage = (
    <Message showIcon type='success' header='Sucesso'>
      Bot criado com sucesso.
    </Message>
  );
  return (
    <>
      <SystemModal
        title='Criar bot'
        open={openCreateBotModal}
        handleClose={() => setCreateBotModal(false)}
        handleSubmit={handleSubmit(onSubmit)}
        isLoading={isLoading}
      >
        <TextField
          className='my-2'
          size='small'
          fullWidth
          placeholder='Nome do bot'
          {...register('botName')}
        />
      </SystemModal>
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
      <BotModal
        openCreateBotModal={openCreateBotModal}
        setCreateBotModal={setCreateBotModal}
      />
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
              botYaml
                ? 'bg-blue-600 text-white my-3'
                : 'bg-gray-500 text-white my-3'
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
