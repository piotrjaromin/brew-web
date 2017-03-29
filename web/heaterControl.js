'use strict';

const React = require('react');
const axios = require('axios');

const Row = require('react-bootstrap/lib/Row');
const Col = require('react-bootstrap/lib/Col');
const ToggleButton = require('react-toggle-button');


class HeaterControl extends React.Component {

    constructor(params) {
        super(params);
        this.state = {heaters: {}};
        this.toggleHeater = this.toggleHeater.bind(this);
        this.heaterControl = this.heaterControl.bind(this);
        this.heaterState = this.heaterState.bind(this);
        this.updateHeatersState = this.updateHeatersState.bind(this)

        setInterval(this.updateHeatersState, 3000)
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
        axios.get(`/heaters/${heaterNo}`)
            .then(resp => {
                if (resp.status == 200) {
                    self.state.heaters[heaterNo] = resp.data.state;
                    self.setState(self.state)
                }
            });
    }

    toggleHeater(heaterNo) {
        const self = this;
        axios.post(`/heaters/${heaterNo}`)
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
            <Col md={12}>
                Heater {heaterNo}:
            </Col>
            <Col md={12}>
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


module.exports = HeaterControl;