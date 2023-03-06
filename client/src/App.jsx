import { useState } from 'react';

function App() {
    const [name, setName] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        console.log('name', name);
        const response = await fetch('http://localhost:3000/users', {
            method: 'POST',
            body: JSON.stringify({ name }),
            headers: {
                'Content-Type': 'application/json',
            },
        });
        const data = await response.json();
        console.log('data', data);
    };

    return (
        <>
            <form onSubmit={handleSubmit}>
                <input
                    type='text'
                    placeholder='Coloca tu nombre'
                    onChange={(e) => setName(e.target.value)}
                />
                <button type='submit'>Guardar</button>
            </form>
        </>
    );
}

export default App;
