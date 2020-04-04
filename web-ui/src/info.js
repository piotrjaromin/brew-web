import React from 'react';

import { Row, Col } from 'react-bootstrap';

import HeaterControl from './heater-control';
import TempControl from './temp-control';
import TempChart from './temp-chart';

const Info = () => {
    return <div>
        <Row>
            <Col md={2} xs={12}>
                <Row>
                    <HeaterControl/>
                </Row>
                <hr/>
                <Row>
                    <TempControl/>
                </Row>
            </Col>
            <Col md={9} xs={12}>
                <TempChart />
            </Col>
        </Row>
    </div>
}

export default Info;