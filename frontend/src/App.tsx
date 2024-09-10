import React from 'react';
import Header from './layout/Header';
import Sidebar from './layout/Sidebar';
import MainContent from './ticketstatus/MainContent';
import Scanner from './scanner/Scanner';

import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.css';




const App: React.FC = () => {
    return (
        <Router>
        <div className="app">
        <Sidebar /> 
            <Routes>
                    <Route path="/scanner" element={<Scanner />} />
                </Routes> 
                <Routes>
                    <Route path="/ticketstatus" element={<MainContent />} />
                </Routes> 
            <Header />
            
        </div>
        </Router>
       
    );
}

export default App;
