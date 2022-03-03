import StepNode from 'components/StepNode';
import React from 'react';
import { Container, Content, Header, Sidebar } from 'rsuite';
import FlowMap from './flow';
import Bar from './bar';
import config from '../../config';
import Panel from './panel';

const nodeTypes = {
  step: StepNode,
};

const Studio = () => {
  return (
    <div>
      <Container style={{ height: '100vh' }}>
        <Bar />
        <Container>
          <FlowMap />
        </Container>
				<Sidebar><Panel /></Sidebar>
      </Container>
    </div>
  );
};

export default Studio;
