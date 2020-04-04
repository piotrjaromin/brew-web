import React, { useState, useEffect } from 'react';

import { Row, Col, DropdownButton, Dropdown } from 'react-bootstrap';
import createKegClient from './services/kegClient';
import { createSimpleLogger } from 'simple-node-logger';

const kegClient = createKegClient();
const logger = createSimpleLogger();

const HeaterControl = () => {
    const powerLevels = [0, 50, 100];

    const [heaterPower, setHeaterPowerValue] = useState(0);

    useEffect(() => {
        kegClient.getHeaterPower().then(setHeaterPowerValue)
    }, []);

    function setHeaterPower(value) {
        return () => {
            logger.info(`clicked heat, setting value to: ${value}`);
            kegClient.setHeaterPower(value)
            setHeaterPowerValue(value);
        }
    }

    return <div>
        <Row>
            <Col md={12}>
                <p>Heater power</p>
                <DropdownButton id="dropdown-item-button" title={`${heaterPower} %`}>
                    {powerLevels.map(powerLevel =>
                        <Dropdown.Item key={`power-${powerLevel}`} onSelect={setHeaterPower(powerLevel)} as="button">{powerLevel}%</Dropdown.Item>)
                    }
                </DropdownButton>
            </Col>
        </Row>
    </div>
}


export default HeaterControl;