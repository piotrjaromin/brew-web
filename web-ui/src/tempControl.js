import React from 'react';

import { Button, FormGroup, FormLabel, FormControl, ButtonToolbar } from 'react-bootstrap';

import createKegClient from './services/kegClient';
import { createSimpleLogger } from 'simple-node-logger';

const kegClient = createKegClient();
const logger = createSimpleLogger();


class TempControl extends React.Component {

    constructor(params) {
        super(params);
        this.state = {};
        this.handleTempChange = this.handleTempChange.bind(this);
        this.handleSetTemp = this.handleSetTemp.bind(this);
        this.handleDisableTemp = this.handleDisableTemp.bind(this);
    }

    handleTempChange(e) {
        logger.info("setting temp to " + e.target.value);
        this.setState({temp : Number(e.target.value)});
    }

    handleSetTemp(e) {
        kegClient.setTemp(this.state.temp);
    }

    handleDisableTemp(e) {
        kegClient.disableTempControl();
    }

    render() {

        return <form>
            <FormGroup controlId="tempValue">
              <FormLabel>Temperature value</FormLabel>
              <FormControl type="number" value={this.state.temp} placeholder="Enter temperature" onChange={this.handleTempChange}/>
            </FormGroup>

            <ButtonToolbar>
                <Button variant="success" onClick={this.handleSetTemp}>Set</Button>
                <Button variant="danger" onClick={this.handleDisableTemp}>Disable</Button>
            </ButtonToolbar>

        </form>
    }
}

export default TempControl;