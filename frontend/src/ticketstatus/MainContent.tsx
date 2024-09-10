import React from 'react';
import './MainContent.css';
import ClearButton from '../layout/ClearButton';
import Table from './table';


const MainContent: React.FC = () => {
    return (
        <div className="main-content">

            <center>
            <br></br>
            <br></br>
            <br></br>
            <br></br>
            <Table />
            <ClearButton />
            </center>
            
        </div>
    );
};

export default MainContent;
