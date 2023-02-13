import Head from 'next/head';
import Link from 'next/link';
import styles from '../styles/Home.module.css';
import Button from 'react-bootstrap/Button';
import 'bootstrap/dist/css/bootstrap.min.css';

async  function getICA() {
  const req = await fetch("/api/getICA");
  const data = await req.json();
  console.log(data);
  document.getElementById("getICA").innerHTML = data;
}

async  function workflow1() {
  const req = await fetch("/api/workflow1");
  const data = await req.json();
  console.log(data);
  document.getElementById("workflow1").innerHTML = data;
}

async  function workflow1_check() {
  const req = await fetch("/api/workflow1_check");
  const data = await req.json();
  console.log(data);
  document.getElementById("workflow1").innerHTML = data;
}

async  function handleClick1() {
  const req = await fetch("/api/query1");
  const data = await req.json();
  console.log(data);
  document.getElementById("result").innerHTML = data;
}

async  function handleClick2() {
  const req = await fetch("/api/query2");
  const data = await req.json();
  console.log(data);
  document.getElementById("result2").innerHTML = data;
}

async  function handleClick3() {
  const req = await fetch("/api/query3");
  const data = await req.json();
  console.log(data);
  document.getElementById("result3").textContent = data;
}
async  function handleClear() {
  document.getElementById("getICA").textContent = "";
  document.getElementById("workflow1").textContent = "";
  document.getElementById("result").textContent = "";
  document.getElementById("result2").textContent = "";
  // document.getElementById("result3").textContent = "";
}

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>CMP Demo </title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <div className="d-grid gap-2 d-md-flex justify-content-md-end">
        <Link href="/"> Home </Link>
        <Link href="/workflow1"> workflow1 </Link>
        <Link href="/workflow2"> workflow2 </Link>
        <Link href="/workflow3"> workflow3 </Link>
        <Link href="/workflow4"> workflow4 </Link>
        <Link href="/workflow5"> workflow5 </Link>
      </div>      
      <p>      </p>

      <div>
      <div className="d-grid gap-2 d-md-flex justify-content-md-end">
        <Button variant="primary" size="sm" onClick={getICA}> ICA Address </Button>
        <Button variant="primary" size="sm" onClick={workflow1}> workflow1 </Button>
        <Button variant="primary" size="sm" onClick={workflow1_check}> workflow1_check </Button>
        <Button variant="primary" size="sm" onClick={handleClick1}> Query1 </Button>
        <Button variant="primary" size="sm" onClick={handleClick2}> Query2 </Button>
        <Button variant="primary" size="sm" onClick={handleClick3}> Query3 </Button>
        <Button variant="primary" size="lg" onClick={handleClear}> Clear </Button>
      </div>
      <p>      </p>
    </div>

      <main>
        <div className={styles.grid}>

          <div className={styles.card1}>
            <h5> ICA Address &rarr;</h5>
            <div id="getICA"></div>
          </div>

        <div className={styles.card1}>
            <h5> workflow1 &rarr;</h5>
            <div id="workflow1"></div>
        </div>

          {/* <a href="http://localhost:3000/api/query1" className={styles.card1}>
          </a> */}
        <div className={styles.card1}>
            <h5>Query Result 1 &rarr;</h5>
            <div id="result"></div>
        </div>

          <div className={styles.card1}>
            <h5>Query Result 2 &rarr;</h5>
            <div id="result2"></div>
          </div>
          

          {/* <a href="http://localhost:3000/api/query3" className={styles.card3}>
            <h3>Query Result 3 &rarr;</h3>
            <div id="result3"></div>
          </a> */}

        </div>
      </main>

      <footer>
        <a
          href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by{' '}
          <img src="/vercel.svg" alt="Vercel" className={styles.logo} />
        </a>
      </footer>

      <style jsx>{`
        main {
          padding: 5rem 0;
          flex: 1;
          display: flex;
          flex-direction: column;
          justify-content: center;
          align-items: center;
        }
        footer {
          width: 100%;
          height: 100px;
          border-top: 1px solid #eaeaea;
          display: flex;
          justify-content: center;
          align-items: center;
        }
        footer img {
          margin-left: 0.5rem;
        }
        footer a {
          display: flex;
          justify-content: center;
          align-items: center;
          text-decoration: none;
          color: inherit;
        }
        code {
          background: #fafafa;
          border-radius: 5px;
          padding: 0.75rem;
          font-size: 1.1rem;
          font-family: Menlo, Monaco, Lucida Console, Liberation Mono,
            DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace;
        }
      `}</style>

      <style jsx global>{`
        html,
        body {
          padding: 0;
          margin: 0;
          font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto,
            Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue,
            sans-serif;
        }
        * {
          box-sizing: border-box;
        }
      `}</style>
    </div>
  )
}