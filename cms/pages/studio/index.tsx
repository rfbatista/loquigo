import React from 'react';
import { Container, Loader } from 'rsuite';
import { useGetBotQuery, useUpdateBotMutation } from 'services/loquiapi';
import YamlEditor from './yaml-editor';


const Studio = () => {
  const botId = 'teste';
  const { data, isFetching, isLoading } = useGetBotQuery(botId);
  if (isFetching || isLoading)
    return <Loader backdrop inverse center content='loading...' vertical />;
  return (
    <div className='grid place-items-center h-screen'>
      <Container
        className='place-items-center w-full h-full'
        style={{ height: '100vh' }}
      >
        <YamlEditor data={data} botId={botId} />
      </Container>
    </div>
  );
};

export default Studio;
