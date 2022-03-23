import { Dashboard } from '@rsuite/icons';
import React from 'react';
import { Sidenav, Nav, Dropdown } from 'rsuite';
import Group from '@rsuite/icons/legacy/Group';
import Magic from '@rsuite/icons/legacy/Magic';
import GearCircle from '@rsuite/icons/legacy/GearCircle';
import EditorPanel from 'components/panel';

const SidePanel = () => {
  const [activeKey, setActiveKey] = React.useState<string | undefined>('1');
  return (
    <div className='flex'>
      <Sidenav
        expanded={false}
        defaultOpenKeys={['3', '4']}
        activeKey={activeKey}
        onSelect={setActiveKey}
        style={{
          height: '100vh',
          borderRight: '1px solid rgba(230, 230, 230, 1)',
        }}
        appearance={'subtle'}
      >
        <Sidenav.Body>
          <Nav>
            <Nav.Item eventKey='1' icon={<Dashboard />}>
              Dashboard
            </Nav.Item>
            <Nav.Item eventKey='2' icon={<Magic />}>
              Criação
            </Nav.Item>
            <Dropdown
              placement='rightStart'
              eventKey='3'
              title='Advanced'
              icon={<Group />}
            >
              <Dropdown.Item eventKey='3-1'>Geo</Dropdown.Item>
              <Dropdown.Item eventKey='3-2'>Devices</Dropdown.Item>
              <Dropdown.Item eventKey='3-3'>Loyalty</Dropdown.Item>
              <Dropdown.Item eventKey='3-4'>Visit Depth</Dropdown.Item>
            </Dropdown>
            <Dropdown
              placement='rightStart'
              eventKey='4'
              title='Settings'
              icon={<GearCircle />}
            >
              <Dropdown.Item eventKey='4-1'>Applications</Dropdown.Item>
              <Dropdown.Item eventKey='4-2'>Channels</Dropdown.Item>
              <Dropdown.Item eventKey='4-3'>Versions</Dropdown.Item>
              <Dropdown.Menu eventKey='4-5' title='Custom Action'>
                <Dropdown.Item eventKey='4-5-1'>Action Name</Dropdown.Item>
                <Dropdown.Item eventKey='4-5-2'>Action Params</Dropdown.Item>
              </Dropdown.Menu>
            </Dropdown>
          </Nav>
        </Sidenav.Body>
      </Sidenav>
      <div className='grow'>
        <EditorPanel />
      </div>
    </div>
  );
};

export default SidePanel;
