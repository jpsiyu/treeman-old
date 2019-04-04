class Timer {
    constructor() {
        this.handler = null
        this.duration = 100
        this.pass = 0
    }

    reset() {
        this.pass = 0
    }

    start(exec) {
        this.reset()
        this.handler = setInterval(
            _ => {
                if (exec) exec()
                this.pass += 0.1
            },
            this.duration
        )
    }

    stop() {
        clearInterval(this.handler)
    }

    getTimePass() {
        return this.pass.toFixed(1)
    }
}

export default Timer