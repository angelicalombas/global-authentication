import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const Register: React.FC = () => {
    const [username, setUsername] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [error, setError] = useState<string>('');
    const navigate = useNavigate();

    const handleRegister = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setError('');

        try {
            await axios.post('http://localhost:8000/register', { username, password });
            alert('Usuário cadastrado com sucesso!');
            navigate('/');
        } catch (error) {
            setError('O username deve ter entre 3 e 20 caracteres e a senha deve ter no mínimo 8 caracteres.');
            console.error(error);
        }
    };

    return (
        <div>
            <h2>Cadastrar Novo Usuário</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <form onSubmit={handleRegister}>
                <input
                    type="text"
                    placeholder="Username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type="submit">Cadastrar</button>
            </form>
        </div>
    );
};

export default Register;