import React, { useState } from 'react';

import { Button, FormGroup, FormLabel, FormControl, ButtonToolbar } from 'react-bootstrap';

import createKegClient from './services/kegClient';
import { createSimpleLogger } from 'simple-node-logger';

const kegClient = createKegClient();
const logger = createSimpleLogger();

const TempControl = () => {
    const [ temp, setTemp ] = useState(0);

    function handleTempChange(e) {
        const val = Number(e.target.value);
        logger.info("setting temp to " + val);
        setTemp(val);
    }

    function handleSetTemp() {
        kegClient.setTemp(temp);
    }

    function handleDisableTemp() {
        kegClient.disableTempControl();
    }

    return <form>
        <FormGroup controlId="tempValue">
            <FormLabel>Temperature value</FormLabel>
            <FormControl type="number" value={temp} placeholder="Enter temperature" onChange={handleTempChange}/>
        </FormGroup>

        <ButtonToolbar>
            <Button variant="success" onClick={handleSetTemp}>Set</Button>
            <Button variant="danger" onClick={handleDisableTemp}>Disable</Button>
        </ButtonToolbar>
    </form>
}

export default TempControl;