import ReactFlow, { Controls } from 'react-flow-renderer';
import { useGetStepQuery } from 'services/loquiapi';
import StepNode from 'components/StepNode';
import { useSelector } from 'react-redux';
import { selectActiveFlow } from 'store/flow';
import { Content, Header } from 'rsuite';
import Panel from './panel';
import { elements } from 'data/nodes';
import React from 'react';
import useGetSize from 'hooks/useGetSize';

const nodeTypes = {
  step: StepNode,
};

const Map = ({ containerRef }) => {
  const [{ width, height }] = useGetSize(containerRef);
  return (
    <div style={{ height: height - 65, width: width }}>
      <ReactFlow elements={elements} nodeTypes={nodeTypes}>
        <Controls />
      </ReactFlow>
    </div>
  );
};

const FlowMap = () => {
  const activeFlow = useSelector(selectActiveFlow);
  const ParentReference = React.useRef(null);

  return (
    <div
      className='relative'
      ref={ParentReference}
      style={{ height: '99%', width: '99%' }}
    >
      <Header>
        <div className='m-4 text-4xl' style={{ height: '40px' }}>
          {activeFlow.name}
        </div>
      </Header>
      <Content>
        <Map containerRef={ParentReference} />
      </Content>
      <div className='absolute right-10 top-10'>
        <Panel />
      </div>
    </div>
  );
};

export default FlowMap;
