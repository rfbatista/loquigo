import ReactFlow, { Background, Controls } from 'react-flow-renderer';
import { useGetFlowMapQuery, useGetStepQuery } from 'services/loquiapi';
import StepNode from 'components/StepNode';
import { useSelector } from 'react-redux';
import { selectActiveFlow } from 'store/flow';
import { Content, Header } from 'rsuite';
import React from 'react';
import useGetSize from 'hooks/useGetSize';

const nodeTypes = {
  step: StepNode,
};

const Map = ({ containerRef, activeFlow }) => {
  const [{ width, height }] = useGetSize(containerRef);
  const { data, isLoading } = useGetFlowMapQuery(activeFlow?.id);
  return (
    <div style={{ height: height - 65, width: width }}>
      <ReactFlow elements={data ?? []} nodeTypes={nodeTypes}>
        <Background color={'#F6F8FA'} />
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
        <div className='relative m-4 text-4xl' style={{ height: '40px' }}>
          {activeFlow.name}
        </div>
      </Header>
      <Content className='relative'>
        <Map containerRef={ParentReference} activeFlow={activeFlow} />
      </Content>
    </div>
  );
};

export default FlowMap;
