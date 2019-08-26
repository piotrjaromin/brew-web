import React from 'react';

import { Row, Col, Button } from 'react-bootstrap';
import { LineChart, Line, CartesianGrid, YAxis, XAxis, Tooltip } from 'recharts';

import axios from 'axios';
import moment from 'moment';

import HeaterControl from './heaterControl';
import TempControl from './tempControl';

import config from './config';
const { backendUrl } = config;


class Info extends React.Component {

    constructor(params) {
        super(params);
        this.refresh = this.refresh.bind(this);
        this.state = {data: []};

        setInterval(this.refresh, 10000);
    }

    refresh() {

        axios.get(`${backendUrl}temperatures`).then(resp => {
            if (resp.status === 200) {
                this.setState({
                    data: resp.data.map(point => {
                        return {
                            value: point.value,
                            timestamp: moment(point.timestamp).format("h:mm:ss")
                        }
                    })
                })
            }
        });
    }

    render() {
        return <div>
            <Row>
                <Col md={2} xs={12}>
                    <Row>
                        <Button variant="info" onClick={this.refresh}>Refresh</Button>
                    </Row>
                    <Row>
                        <HeaterControl/>
                    </Row>
                    <hr/>
                    <Row>
                        <TempControl/>
                    </Row>
                </Col>
                <Col md={9} xs={12}>
                    <Row>
                        <h2 className="text-center">Temperature chart</h2>
                    </Row>
                    <Row>
                        <LineChart width={800} height={400} data={this.state.data}>
                            <XAxis dataKey="timestamp"/>
                            <YAxis />
                            <CartesianGrid strokeDasharray="3 3"/>
                            <Tooltip />
                            <Line type="monotone" dataKey="value" stroke="#8884d8"/>
                        </LineChart>
                    </Row>
                </Col>
            </Row>
        </div>
    }
}

export default Info;