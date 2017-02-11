'use strict';

const React = require('react');
const Tabs = require('react-bootstrap/lib/Tabs');
const Tab = require('react-bootstrap/lib/Tab');
const Row = require('react-bootstrap/lib/Row');
const Col = require('react-bootstrap/lib/Col');
const PageHeader = require('react-bootstrap/lib/PageHeader');
const ReactDOM = require('react-dom');

require('bootstrap/dist/css/bootstrap.css');

//Content
const Info = require('./info');

class App extends React.Component {

    constructor(params) {
        super(params)
    }

    render() {
        return <div className="container">
            <h1 className="text-center">
                Home brew
            </h1>
            <Tabs defaultActiveKey={1} animation={false} id="noanim-tab-example">
                <Tab eventKey={1} title="Info"><Info/></Tab>
            </Tabs>
        </div>
    }
}

ReactDOM.render(<App/>, document.getElementById('app'));