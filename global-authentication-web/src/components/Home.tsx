import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Home: React.FC = () => {
    const [message, setMessage] = useState<string>('');
    const [error, setError] = useState<string>('');

    useEffect(() => {
        const fetchMessage = async () => {
            try {
                const token = localStorage.getItem('token');
                const response = await axios.get('http://localhost:8000/home', {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                setMessage(response.data.message);
            } catch (error) {
                setError('Você não tem permissão para acessar o conteúdo desta página');
                console.error(error);
            }
        };

        fetchMessage();
    }, []);

    return (
        <div>
            <h2>Home</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <p>{message}</p>
        </div>
    );
};

export default Home;