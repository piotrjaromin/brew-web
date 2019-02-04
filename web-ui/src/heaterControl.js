import React from 'react';

import { Row, Col } from 'react-bootstrap';
import ToggleButton from 'react-toggle-button';
import createKegClient from './services/kegClient';

const kegClient = createKegClient();

const HEATER_1 = '1';
const HEATER_2 = '2';

class HeaterControl extends React.Component {

    constructor(params) {
        super(params);
        this.state = {
            heaters: {
                [HEATER_1]: false,
                [HEATER_2]: false,
            }
        };

        this.toggleHeater = this.toggleHeater.bind(this);
        this.heaterControl = this.heaterControl.bind(this);
        this.heaterState = this.heaterState.bind(this);
        this.updateHeatersState = this.updateHeatersState.bind(this);
        this.toggleHeater = this.toggleHeater.bind(this);
    }

    updateHeatersState() {
        Promise.all([
            this.heaterState(HEATER_1),
            this.heaterState(HEATER_2),
        ])
            .then(([state1, state2]) => {
                console.log(`setting state for \n${HEATER_1}: ${state1}\n${HEATER_2}: ${state2}`);
                this.setState(prevState => ({
                    heaters: {
                        ...prevState.heaters,
                        [HEATER_1]: state1,
                        [HEATER_2]: state2,
                    }
                }));
            })
    }

    componentDidMount() {
        this.updateHeatersState()
    }


    heaterState(heaterNo) {
        return kegClient.getHeaterState(heaterNo);
    }

    toggleHeater(heaterNo) {
        kegClient.toggleHeater(heaterNo)
            .then(this.updateHeatersState());
    }

    heaterControl(heaterNo) {
        const self = this;
        return <Row>
            <Col md={5}>
                Heater{heaterNo}:
            </Col>
            <Col md={7}>
                <ToggleButton
                    value={ self.state.heaters[heaterNo] }
                    onToggle={() => self.toggleHeater(heaterNo)}/>
            </Col>
        </Row>
    }

    render() {

        return <div>
            <Row>
                <Col md={12}>
                    {this.heaterControl(HEATER_1)}
                </Col>
            </Row>
            <Row>
                <Col md={12}>
                    {this.heaterControl(HEATER_2)}
                </Col></Row>
        </div>
    }
}


export default HeaterControl;