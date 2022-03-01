import React from 'react';
import { Sidebar, Sidenav, Nav, Dropdown } from 'rsuite';
import { Dashboard } from '@rsuite/icons';
import { useGetFlowQuery } from 'services/loquiapi';
import { setActiveFlow } from 'store/flow';
import { useDispatch } from 'react-redux';

const FlowDropDown = ({ item }) => {
  const dispatch = useDispatch();

  return (
    <>
      <Dropdown.Item
        eventKey={item.id}
        onClick={() => {
          dispatch(setActiveFlow(item));
        }}
      >
        {item.name}
      </Dropdown.Item>
    </>
  );
};

const Bar = () => {
  const [expanded, setExpanded] = React.useState(true);
  const [activeKey, setActiveKey] = React.useState('1');
  const [expand, setExpand] = React.useState(true);
  const { data, isError, isLoading, error } = useGetFlowQuery('Welcome BOT');

  return (
    <Sidebar
      style={{ display: 'flex', flexDirection: 'column', height: '100vh' }}
      width={expand ? 260 : 56}
      collapsible
    >
      <Sidenav
        expanded={expanded}
        defaultOpenKeys={['3', '4']}
        activeKey={activeKey}
        onSelect={(key, _) => setActiveKey(String(key))}
        style={{ height: '100vh' }}
      >
        <Sidenav.Body>
          <Nav>
            <Dropdown
              placement='rightStart'
              eventKey='3'
              title='Flows'
              icon={<Dashboard />}
            >
              {isLoading ? (
                <></>
              ) : (
                data.map((item, idx) => (
                  <FlowDropDown key={item.id} item={item} />
                ))
              )}
            </Dropdown>
          </Nav>
        </Sidenav.Body>
      </Sidenav>
    </Sidebar>
  );
};

export default Bar;
