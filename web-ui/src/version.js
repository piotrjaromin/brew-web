import React, { useState } from 'react';
import axios from 'axios';

import config from './config';

const { backendUrl } = config;

const Version = () => {
    const [version, setVersion] = useState(0);

    axios.get(`${backendUrl}version`).then(resp => {
        if (resp.status === 200) {
            setVersion(resp.data.version);
        }
    })

    return <div>
        {version}
    </div>
}


export default Version;