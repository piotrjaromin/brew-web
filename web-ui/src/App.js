import React from 'react';
import { Tabs, Tab } from 'react-bootstrap';

import 'bootstrap/dist/css/bootstrap.css';
import './App.css';

//Content
import Info from './info';

class App extends React.Component {

    render() {
        return <div className="container">
            <h1 className="text-center">
                Home brew
            </h1>
            <Tabs defaultActiveKey={1} animation="false" id="noanim-tab-example">
                <Tab eventKey={1} title="Info"><Info/></Tab>
            </Tabs>
        </div>
    }
}

export default App;