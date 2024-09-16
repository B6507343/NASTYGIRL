import React from 'react';
import Header from './components/Header';
import Sidebar from './components/Sidebar';
import Table from './ticketstatus/ticketstatus';
import QrScanner from './Qrscanner/QrScanner';

import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.css';





const App: React.FC = () => {
    return (
        <Router>
        <div className="app">
        <Sidebar /> 
            <Routes>
                    <Route path="/scanner" element={<QrScanner />} />
                </Routes> 
                <Routes>
                    <Route path="/ticketstatus" element={<Table />} />
                </Routes> 
            <Header />
            
        </div>
        </Router>
       
    );
}

export default App;
