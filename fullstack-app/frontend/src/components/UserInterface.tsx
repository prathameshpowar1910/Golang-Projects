import React from 'react'
import axios from 'axios'
import CardComponent from './CardComponent'

interface User {
  id: number;
  name: string;
  email: string;
}

interface UserInterfaceProps {
  backendName: string; 
}


export const UserInterface = () => {
  return (
    <div>UserInterface</div>
  )
}
