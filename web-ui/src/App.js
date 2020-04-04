import React from 'react';
import { Tabs, Tab, Row, Col } from 'react-bootstrap';

import 'bootstrap/dist/css/bootstrap.css';
import './app.css';

//Content
import Info from './info';
import Version from './version';

const App = () => {
    return <div className="container">
        <Row>
            <Col xs={11}></Col>
            <Col xs={1}><Version/></Col>
        </Row>
        <Row>
            <Col xs={4}></Col>
            <Col xs={4}>
                <h1 className="text-center">
                    Home brew
                </h1>
            </Col>
            <Col xs={4}></Col>
        </Row>

        <Tabs defaultActiveKey={1} animation="false" id="noanim-tab-example">
            <Tab eventKey={1} title="Info"><Info/></Tab>
        </Tabs>
    </div>
}

export default App;