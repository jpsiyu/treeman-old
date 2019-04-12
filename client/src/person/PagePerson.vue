<template>
    <div class="pagePerson">
        <Banner @pageSwitch="pageSwitch"></Banner>
        <div class="pagePerson-search">
            <input
                type="text"
                placeholder="search name"
                @keydown.enter="search"
                v-model="searchName"
            >
            <button @click="search">Search</button>
        </div>
        <div class="pagePerson-list">
            <Member
                v-for="info in members"
                :key="info.id"
                :info="info"
                @getAllPerson="getAllPerson"
                @pageSwitch="pageSwitch"
            />
        </div>
        <PageAddPerson
            v-if="pageAddState"
            :pageAddData="pageAddData"
            @pageSwitch="pageSwitch"
            @getAllPerson="getAllPerson"
        ></PageAddPerson>
    </div>
</template>

<script>
import { mapState } from "vuex";
import request from "../request";
import shuffle from "../lib/shuffle";
import Member from "./Member";
import Banner from "../common/Banner";
import PageAddPerson from "./PageAddPerson";
export default {
    name: "PagePerson",
    data: function() {
        return {
            pageAddState: false,
            pageAddData: null,
            searchName: ""
        };
    },
    components: { Member, Banner, PageAddPerson },
    computed: {
        ...mapState({
            members: state => state.person.members
        })
    },
    watch: {
        searchName: function() {
            this.search();
        }
    },
    methods: {
        pageSwitch: function(b, data) {
            this.pageAddState = b;
            if (data) this.pageAddData = data;
            if (!b) this.pageAddData = null;
        },
        getAllPerson: function() {
            request
                .getAllPerson()
                .then(res => {
                    const serverData = res.data;
                    const data = serverData.data || [];
                    shuffle(data);
                    this.$store.commit("initMembers", data);
                })
                .catch(_ => {});
        },
        search: function() {
            if (this.searchName == "") {
                this.getAllPerson();
                return;
            }
            request.getPersonByName(this.searchName).then(res => {
                const serverData = res.data;
                const data = serverData.data || [];
                this.$store.commit("initMembers", data);
            });
        }
    },
    mounted: function() {
        this.getAllPerson();
    }
};
</script>

<style scoped>
.pagePerson {
    position: relative;
}

.pagePerson-search {
    max-width: 500px;
    margin: 30px auto;
    display: flex;
    align-items: flex-end;
}

.pagePerson-search input {
    width: 60%;
    border-radius: 10px;
    outline: none;
    height: 30px;
    border: none;
    font-size: 16px;
    padding: 3px 10px;
    box-sizing: border-box;
}

.pagePerson-search button {
    outline: none;
    border: none;
    background-color: seagreen;
    color: white;
    padding: 0 10px;
    border-radius: 10px;
    cursor: pointer;
    height: 25px;
    margin: 0 10px;
}

.pagePerson-list {
    display: flex;
    margin: auto;
    max-width: 500px;
    justify-content: center;
    flex-direction: column;
    padding: 10px;
}
</style>
