import React from 'react';

import { Row } from 'react-bootstrap';
import { LineChart, Line, CartesianGrid, YAxis, XAxis, Tooltip } from 'recharts';

import axios from 'axios';
import moment from 'moment';

import config from './config';

const { backendUrl } = config;

const TEMP_REFRESH_INTERVAL = 5000;

class TempChart extends React.Component {
    state = {
        temperatures: [],
    };

    constructor(props) {
        super(props);

        this.refresh = this.refresh.bind(this);
        this.interval = setInterval(this.refresh, TEMP_REFRESH_INTERVAL);
    }


    refresh() {
        axios.get(`${backendUrl}temperatures`).then(resp => {
            if (resp.status === 200) {
                const temperatures = resp.data.map(point => ({
                    value: point.value,
                    timestamp: moment(point.timestamp).format("h:mm:ss")
                }));

                this.setState({
                    temperatures
                })
            }
        });
    }

    componentDidMount() {
        this.refresh()
    }

    componentWillUnmount() {
        clearInterval(this.interval);
    }

    render() {
        return <div>
            <Row>
                <h2 className="text-center">Temperature chart</h2>
            </Row>
            <Row>
                <LineChart width={800} height={400} data={this.state.temperatures}>
                    <XAxis dataKey="timestamp"/>
                    <YAxis />
                    <CartesianGrid strokeDasharray="3 3"/>
                    <Tooltip />
                    <Line type="monotone" dataKey="value" stroke="#8884d8"/>
                </LineChart>
            </Row>
        </div>
    }
}

export default TempChart;