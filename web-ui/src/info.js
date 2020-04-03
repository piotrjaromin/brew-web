import React, { useState } from 'react';

import { Row, Col, Button } from 'react-bootstrap';
import { LineChart, Line, CartesianGrid, YAxis, XAxis, Tooltip } from 'recharts';

import axios from 'axios';
import moment from 'moment';

import HeaterControl from './heaterControl';
import TempControl from './tempControl';

import config from './config';

const { backendUrl } = config;

const TEMP_REFRESH_INTERVAL = 5000;

const Info = () => {
    const [ temperatures, setTemperatures ] = useState(0);

    setInterval(refresh, TEMP_REFRESH_INTERVAL);

    function refresh() {
        axios.get(`${backendUrl}temperatures`).then(resp => {
            if (resp.status === 200) {
                const temps = resp.data.map(point => ({
                    value: point.value,
                    timestamp: moment(point.timestamp).format("h:mm:ss")
                }));

                setTemperatures(temps);
            }
        });
    }

    return <div>
        <Row>
            <Col md={2} xs={12}>
                <Row>
                    <Button variant="info" onClick={refresh}>Refresh</Button>
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
                    <LineChart width={800} height={400} data={temperatures}>
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

export default Info;