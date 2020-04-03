import axios from 'axios';
import { createSimpleLogger } from 'simple-node-logger';
import config from '../config';

const { backendUrl } = config;

const logger = createSimpleLogger();

export default function create() {
    logger.info(`Backend url is ${backendUrl}`);

    function setHeaterPower(power) {
        logger.info(`setting heater to power ${power}`);
        return axios.post(`${backendUrl}heaters`, { power })
            .then(resp => {
                if (resp.status !== 200) {
                    return logger.error(`Invalid response from backend. ${resp.status}: ${resp.data}`);
                }

                return resp.data.state;
            })
            .catch(logger.error);
    }

    function getHeaterPower() {
        return axios.get(`${backendUrl}heaters`)
            .then(resp => {
                if (resp.status === 200) {
                    return resp.data.power;
                }
            })
            .catch(logger.error);
    }

    function setTemp(temp) {
        axios
            .post(`${backendUrl}temperatures/control`, { value: temp})
            .then( resp => {
                if (resp.status !== 200) {
                    return logger.error(`Wrong status code response for set temp ${resp.statusCode}`);
                }

                return logger.info("temp was set successfully: " + temp);
            })
            .catch(e => logger.error(`Error while setting temperature to control. ${e}`))
    }

    function disableTempControl() {
        axios
            .delete(`${backendUrl}temperatures/control`)
            .then( resp => {
                if (resp.status !== 200) {
                    return logger.error(`Wrong status code for delete temp response ${resp.statusCode}`);
                }

                return logger.info('temp was deleted successfully');
            })
            .catch(e => logger.error(`Error while deleting temperature to control. ${e}` ))
    }

    return {
        setHeaterPower,
        getHeaterPower,
        setTemp,
        disableTempControl
    }
}