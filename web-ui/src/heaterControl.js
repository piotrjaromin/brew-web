import React, { useState } from 'react';

import { Row, Col } from 'react-bootstrap';
import createKegClient from './services/kegClient';
import { createSimpleLogger } from 'simple-node-logger';

import Slider from 'react-rangeslider'

import 'react-rangeslider/lib/index.css'

const kegClient = createKegClient();
const logger = createSimpleLogger();

const HEATER_REFRESH_INTERVAL = 5000;

const HeaterControl = () => {
    const [heaterPower, setHeaterPowerValue] = useState(0);

    setInterval(updateHeatersState, HEATER_REFRESH_INTERVAL);

    function updateHeatersState() {
        kegClient.getHeaterPower().then(setHeaterPowerValue)
    }

    function setHeaterPower(value) {
        logger.info(`clicked heat, setting value to: ${value}`);
        kegClient.setHeaterPower(value)
        setHeaterPowerValue(value);
    }

    return <div>
        <Row>
            <Col md={12}>
                <p>Heater power</p>
                <Slider
                    min={0}
                    max={100}
                    step={50}
                    value={heaterPower}
                    onChange={setHeaterPower}
                />
            </Col>
        </Row>
    </div>
}


export default HeaterControl;