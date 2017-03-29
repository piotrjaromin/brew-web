'use strict';

const React = require('react');
const axios = require('axios');


const moment = require('moment');

const Row = require('react-bootstrap/lib/Row');
const Col = require('react-bootstrap/lib/Col');
const Button = require('react-bootstrap/lib/Button');

const LineChart = require('recharts/lib').LineChart;
const Line = require('recharts/lib').Line;
const CartesianGrid = require('recharts/lib').CartesianGrid;
const YAxis = require('recharts/lib').YAxis;
const XAxis = require('recharts/lib').XAxis;
const Tooltip = require('recharts/lib').Tooltip;

const HeaterControl = require('./heaterControl');
const TempControl = require('./tempControl');

class Info extends React.Component {

    constructor(params) {
        super(params);
        this.refresh = this.refresh.bind(this);
        this.state = {data: []};

        setInterval(this.refresh, 10000);
    }

    refresh() {

        axios.get('temperatures').then(resp => {
            if (resp.status == 200) {
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
                        <Button bsStyle="info" onClick={this.refresh}>Refresh</Button>
                    </Row>
                    <Row>
                        <HeaterControl/>
                    </Row>
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


module.exports = Info;