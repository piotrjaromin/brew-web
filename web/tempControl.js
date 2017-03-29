'use strict';

const React = require('react');
const axios = require('axios');

const Col = require('react-bootstrap/lib/Col');
const Button = require('react-bootstrap/lib/Button');
const FormGroup = require('react-bootstrap/lib/FormGroup');
const ControlLabel = require('react-bootstrap/lib/ControlLabel');
const FormControl = require('react-bootstrap/lib/FormControl');

class TempControl extends React.Component {

    constructor(params) {
        super(params);
        this.state = {};
        this.handleTempChange = this.handleTempChange.bind(this);
        this.handleSetTemp = this.handleSetTemp.bind(this);
        this.handleDisableTemp = this.handleDisableTemp.bind(this);
    }

    handleTempChange(e) {
        console.log(e)
        console.log("setting temp to " + e.target.value)
        this.setState({temp : Number(e.target.value)})
    }

    handleSetTemp(e) {
        const self = this
        axios
        .post(`/temperatures/control`, { value: this.state.temp})
        .then( resp => {
             if (resp.status === 200) {
                console.log("temp was set successfully: " + self.state.temp)
            } else {
                console.log(`Wrong status code response for set temp ${resp.statusCode}`)
            }   
        } )
        .catch(e => {
            console.log(`Error while setting temperature to control. ${e}` )
        })
    
    }

    handleDisableTemp(e) {
        axios
        .delete(`/temperatures/control`)
        .then( resp => {
             if (resp.status === 200) {
                console.log("temp was deleted successfully")
            } else {
                console.log(`Wrong status code for delete temp response ${resp.statusCode}`)
            }   
        } )
        .catch(e => {
            console.log(`Error while deleting temperature to control. ${e}` )
        })
    } 

    render() {

        return <form>
            <FormGroup controlId="tempValue">
              <ControlLabel>Temeprature value</ControlLabel>
              <FormControl type="number" value={this.state.temp} placeholder="Enter temperature" onChange={this.handleTempChange}/>
            </FormGroup>

            <Button bsStyle="success" onClick={this.handleSetTemp}>Set</Button>
            <Button bsStyle="danger" onClick={this.handleDisableTemp}>Disable</Button>

        </form>
    }
}

module.exports = TempControl