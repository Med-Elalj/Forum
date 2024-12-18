const url  = 'https://api.coindesk.com/v1/bpi/currentprice.json'

async function x(){
    try{
        const respond = await fetch(url)
        console.log(respond);
        const result = await respond.json()
        console.log(result);
        
    }catch{
        console.log("ERorrldkfj ");

    }
}

x()