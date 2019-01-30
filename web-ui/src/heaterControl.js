import React from 'react';
import axios from 'axios';

import { Row, Col } from 'react-bootstrap';
import ToggleButton from 'react-toggle-button';

import config from './config';
const { backendUrl } = config;

class HeaterControl extends React.Component {

    constructor(params) {
        super(params);
        this.state = {heaters: {}};
        this.toggleHeater = this.toggleHeater.bind(this);
        this.heaterControl = this.heaterControl.bind(this);
        this.heaterState = this.heaterState.bind(this);
        this.updateHeatersState = this.updateHeatersState.bind(this);

        setInterval(this.updateHeatersState, 15000)
    }

    updateHeatersState() {
        this.heaterState(1);
        this.heaterState(2);
    }

    componentDidMount() {
        this.updateHeatersState()
    }


    heaterState(heaterNo) {
        const self = this;
        axios.get(`${backendUrl}/heaters/${heaterNo}`)
            .then(resp => {
                if (resp.status === 200) {
                    self.state.heaters[heaterNo] = resp.data.state;
                    self.setState(self.state)
                }
            });
    }

    toggleHeater(heaterNo) {
        const self = this;
        axios.post(`${backendUrl}/heaters/${heaterNo}`)
            .then(resp => {
                if (resp.status === 200) {
                    self.state.heaters[heaterNo] = resp.data.state;
                    self.setState(self.state)
                } else {
                    console.error(`Invalid response from backend. ${resp.status}: ${resp.data}`)
                }
            });
    }

    heaterControl(heaterNo) {
        const self = this;
        return <Row>
            <Col md={5}>
                Heater{heaterNo}:
            </Col>
            <Col md={7}>
                <ToggleButton
                    value={ self.state.heaters[heaterNo] || false }
                    onToggle={() => self.toggleHeater(heaterNo)}/>
            </Col>
        </Row>
    }

    render() {

        return <div>
            <Row>
                <Col md={12}>
                    {this.heaterControl(1)}
                </Col>
            </Row>
            <Row>
                <Col md={12}>
                    {this.heaterControl(2)}
                </Col></Row>
        </div>
    }
}


export default HeaterControl;