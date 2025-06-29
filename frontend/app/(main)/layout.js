'use client'
import React, { use, useEffect } from 'react'
import Header from './_components/header'
import Navigation from './_components/navigation'
import { fetchTestData, test } from '../actions/test'





export default function MainLayout({ children }) {

    useEffect( async()=> {
      let response = await fetch("http://localhost:8080/api/test")
      let data = await response.json()
      console.log(data)
    },[])

  return (
    <>
      <Header />
      <Navigation />
      {children}
    </>
  )
}
