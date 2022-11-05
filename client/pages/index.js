import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'

export default function Home() {

  const makeReq = () => {
    fetch('https://api.first.org/data/v1/epss?cve=CVE-2022-27225', {
      method: 'GET',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': 'strict-origin-when-cross-origin'
      }
    })
  }

  return (
    <div>
      <h1 className="text-3xl font-bold underline">
        Hello world!
      </h1>
      <h2>
        from EPSS Browser
      </h2>

      <button onClick={makeReq}>Call API</button>
    </div>
  )
}
