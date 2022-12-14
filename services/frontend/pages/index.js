import Head from 'next/head'
import styles from '../styles/Home.module.css'
import {useEffect, useState} from "react";

export async function getServerSideProps() {
    const initialData = await fetch("http://localhost:8000/handler-initial-data").then(x => x.json());
    return {props: {data: initialData}}
}

export default function Home(props) {
    const [data, setData] = useState(props.data);
    const [ws, setWS] = useState(null);
    useEffect(() => {
        const newWS = new WebSocket("ws://localhost:8000/handler")
        newWS.onerror = err => console.error(err);
        newWS.onopen = () => setWS(newWS);
        newWS.onmessage = msg => setData(JSON.parse(msg.data));
    }, [])
    return (
        <div className={styles.container}>
            <Head>
                <title>OSS Docs</title>
                <meta name="description" content="Like Google Docs, but open source!"/>
                <link rel="icon" href="/favicon.ico"/>
            </Head>

            <main className={styles.main}>
                <h1 className={styles.title}>
                    {props.title || "Untitled Document"}
                </h1>
                <div>Data is: {JSON.stringify(data)}</div>
            </main>
        </div>
    )
}
