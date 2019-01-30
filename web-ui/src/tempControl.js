import React from 'react';
import axios from 'axios';

import { Button, FormGroup, FormLabel, FormControl } from 'react-bootstrap';

class TempControl extends React.Component {

    constructor(params) {
        super(params);
        this.state = {};
        this.handleTempChange = this.handleTempChange.bind(this);
        this.handleSetTemp = this.handleSetTemp.bind(this);
        this.handleDisableTemp = this.handleDisableTemp.bind(this);
    }

    handleTempChange(e) {
        console.log(e);
        console.log("setting temp to " + e.target.value);
        this.setState({temp : Number(e.target.value)});
    }

    handleSetTemp(e) {
        const self = this;
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
            console.log(`Error while setting temperature to control. ${e}` );
        })

    }

    handleDisableTemp(e) {
        axios
        .delete(`/temperatures/control`)
        .then( resp => {
             if (resp.status === 200) {
                console.log("temp was deleted successfully");
            } else {
                console.log(`Wrong status code for delete temp response ${resp.statusCode}`);
            }
        } )
        .catch(e => {
            console.log(`Error while deleting temperature to control. ${e}` );
        })
    }

    render() {

        return <form>
            <FormGroup controlId="tempValue">
              <FormLabel>Temeprature value</FormLabel>
              <FormControl type="number" value={this.state.temp} placeholder="Enter temperature" onChange={this.handleTempChange}/>
            </FormGroup>

            <Button bsStyle="success" onClick={this.handleSetTemp}>Set</Button>
            <Button bsStyle="danger" onClick={this.handleDisableTemp}>Disable</Button>

        </form>
    }
}

export default TempControl;