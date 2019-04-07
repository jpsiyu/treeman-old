<template>
    <div class="miss">
        <h2>Where, where, where am i?</h2>
        <img src="../assets/images/man.png">
    </div>
</template>

<script>
import axios from 'axios'
export default {
    mounted: function(){
       axios.get("/pubapi/test") 
        .then(res => {
            console.log(res)
        })
        .catch(err => {
            console.log(err)
        })
    }
};
</script>

<style scoped>
.miss {
    display: flex;
    flex-direction: column;
    margin: auto;
    max-width: 400px;
    align-items: center;
}

.miss img{
    width: 100%;
    height: auto;
}
</style>
