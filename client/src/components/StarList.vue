<template>
<div>
    <div class="row justify-content-center">
        <div class="col-md-6 list">
            <form v-on:submit.prevent @submit="searchName">
                <div class="input-group mb-3">
                    <input ref="name" type="text" class="form-control fs-20" placeholder="Search by name" v-model="name" />
                    <div class="input-group-append">
                        <button class="btn btn-outline-primary ml-1 fs-20" type="submit">
                            Search
                        </button>
                        <button class="btn btn-outline-secondary fs-20" type="button" @click="refreshList">
                            Reset
                        </button>
                    </div>
                </div>
            </form>
        </div>
    </div>
    <div class="row justify-content-center">
        <div class="col-md-6">
            <h4 style="font-size: 35px;">Available Stars</h4>
            <div class="btn-group dropleft mb-2">
                <button type="button" class="btn btn-secondary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    Star Class
                </button>
                <div class="dropdown-menu">
                    <a class="dropdown-item" v-bind:href="'/stars?page=1&class=All'" :class="{'active':(starClass === 'All')}">All</a>
                    <a class="dropdown-item" v-bind:href="'/stars?page=1&class=Yellow Dwarf'" :class="{'active':(starClass === 'Yellow Dwarf')}">Yellow Dwarf</a>
                    <a class="dropdown-item" v-bind:href="'/stars?page=1&class=White Dwarf'" :class="{'active':(starClass === 'White Dwarf')}">White Dwarf</a>
                </div>
            </div>
            <transition-group name="slide-fade" tag="ul" class="list-group">
                <li class="list-group-item lgi-pointer py-1" :class="{ active: index == currentIndex }" v-for="(star, index) in stars" :key="star.star_id" @click="setActiveStar(star, index)">
                    {{ star.name }}
                </li>
            </transition-group>
            <nav aria-label="Pagination">
                <ul class="pagination mt-2 justify-content-center">
                    <li v-if="loaded" :key="'prev'" :class="{'disabled':currentPage === 1}" class="page-item previous-item">
                        <a class="page-link" v-bind:href="'/stars?page='+ (currentPage-1) + '&class=' + starClass">Prev</a>
                    </li>
                    <li v-for="page in pageArray" :key="page" class="page-item" :class="{'active':(currentPage === page)}">
                        <a class="page-link" v-bind:href="'/stars?page='+ page + '&class=' + starClass">{{page}}</a>
                    </li>
                    <li v-if="loaded" :key="'next'" :class="{'disabled':currentPage === totalPages}" class="page-item next-item">
                        <a class="page-link" v-bind:href="'/stars?page='+ (currentPage+1) + '&class=' + starClass">Next</a>
                    </li>
                </ul>
            </nav>
        </div>
    </div>
    <h1 style="text-align: center; margin-top: 2%; font-size: 50px;" v-if="currentStar">{{ currentStar.name }}</h1>
    <div class="row">
        <transition name="slide-fade">
            <div class="col-md-4" v-if="currentStar">
                <h1>Basic Information</h1>
                <ul class="list-group">
                    <li :key="'class'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Class:</strong> {{currentStar.basic_information.class}}</p>
                    </li>
                    <li :key="'mass'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Mass:</strong> {{currentStar.basic_information.mass}}</p>
                    </li>
                    <li :key="'radius'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Radius:</strong> {{currentStar.basic_information.radius}}</p>
                    </li>
                    <li :key="'system'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>System:</strong> {{currentStar.basic_information.system}}</p>
                    </li>
                    <li :key="'temperature'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Temperature:</strong> {{currentStar.basic_information.temperature}}</p>
                    </li>
                    <li :key="'age'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Age:</strong> {{currentStar.basic_information.age}}m/s<sup>2</sup></p>
                    </li>
                </ul>
            </div>
        </transition>
        <transition name="slide-fade">
            <img class="img-center col-md-4 justify-content-right" v-if="currentStar" v-bind:src="currentStar.image">
        </transition>
        <div class="col-md-4">
        </div>
    </div>
    <div class="row mt-3 justify-content-center">
        <div class="col-md-10">
            <div v-if="currentStar">
                <ul class="list-group list-group-flush">
                    <li class="list-group-item mt-2" v-for="(val, index) in currentStar.facts" :key="index">
                        <h1 style="font-size: 40px; text-align: center;">{{val.title}}</h1>
                        <p style="font-size: 22px;">{{val.fact}}</p>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>
</template>

<script>
import StarService from "../services/StarService";

export default {
    name: "stars-list",
    data() {
        return {
            stars: [],
            currentStar: null,
            currentIndex: -1,
            name: "",
            pageArray: [],
            currentPage: parseInt(this.$route.query.page),
            totalPages: 0,
            starClass: this.$route.query.class,
            loaded: false
        };
    },
    methods: {
        retrieveStars(page, starClass) { // Fetchs all of our stars for the current page.
            StarService.getAll(page, starClass)
                .then(response => {
                    this.stars = response.data.stars;
                    console.log(response.data);
                    console.log(this.pageArray);
                    this.currentPage = page;
                    this.pageArray = this.generatePaginationPageArray(response.data.number_of_documents, 5);
                    this.currentStar = null;
                    this.currentIndex = -1;
                })
                .catch(e => {
                    console.log(e);
                });
        },

        refreshList() { // Refreshes the page to the default state.
            this.retrieveStars(this.currentPage, this.starClass);
            this.currentStar = null;
            this.currentIndex = -1;
            this.name = "";
            document.getElementsByClassName("pagination")[0].style.visibility = "visible";
        },

        setActiveStar(star, index) { // Updates currently viewed star.
            this.currentStar = star;
            this.currentIndex = index;
        },

        searchName() { // This will search for a star by name and return it, caps specific as of now.
            if (this.name === "") {
                alert("Please enter a star to search for.");
                this.$refs.name.focus();
                return;
            }

            StarService.get(this.name)
                .then(response => {
                    this.stars = [];
                    this.stars.push(response.data);
                    this.currentStar = null;
                    this.currentIndex = -1;
                    document.getElementsByClassName("pagination")[0].style.visibility = "hidden";
                })
                .catch(e => {
                    if (e.response.status == 404) {
                        alert("No star by that name exists in the database.");
                        this.name = "";
                        this.$refs.name.focus();
                    }
                });
        },
        validatePageCount(count, total) {
            if (count > total) {
                count = total;
            }
            if (count % 2 == 0) { //ensure odd number
                count--;
            }
            return count;
        },
        generatePaginationPageArray(numOfDocuments, recordsPerPage) { // Fetch our pages in an array so we can iterate over it later to create the pagination list items.
            if (numOfDocuments == 0) { // No paging if no records.
                return;
            }

            const pageArray = [];
            this.totalPages = Math.ceil(numOfDocuments / recordsPerPage); // sets total page count for data
            let pageDisplayCount = this.validatePageCount(5, this.totalPages);
            const mid = Math.ceil(pageDisplayCount / 2); // sets mid of displayed pages
            let startIndex = (this.currentPage + 1) - mid; // set first page to be displayed

            if (this.totalPages <= recordsPerPage) {
                startIndex = 1;
                pageDisplayCount = this.totalPages;
            } else if (this.currentPage < mid) { // account for front half
                startIndex = 1;
            } else if ((this.currentPage + (mid - 1)) > this.totalPages) { //account for back half
                startIndex = this.totalPages - pageDisplayCount + 1;
            }
            for (let i = startIndex; i < startIndex + pageDisplayCount; i++) {
                pageArray.push(i);
            }
            return pageArray;
        }
    },
    watch: { // Watch for data change in which page the user is currently on, call API to get new data when it changes.
        // '$route.query.page': {
        //     immediate: true,
        //     handler(page) {
        //         page = parseInt(page) || 1;
        //         if (page !== this.currentPage) {
        //             this.retrieveStars(page, this.type);
        //         }
        //     }
        // },
        // "currentPage": function () {
        //     this.refreshList();
        // },

        // "type": function () {
        //      if (this.currentPage === 1) {   // Otherwise changing current page will take care of the refresh for us.
        //         this.refreshList();
        //     }
        //     this.currentPage = 1;
        // }
    },
    mounted() {
        this.retrieveStars(this.currentPage, this.starClass);
        this.loaded = true;
    }
};
</script>

<style>
.lgi-pointer {
    cursor: pointer;
}

.lgi-pointer:hover {
    background-color: #53a5fc;
}

.lgi-colored-space {
    background-color: #fcfbfe;
    color: #1d1135;
}

.img-center {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 50%;
}

.fs-20 {
    font-size: 20px;
}

/* Enter and leave animations can use different */
/* durations and timing functions.              */
.slide-fade-enter-active {
    transition: all 2s ease;
}

.slide-fade-leave-active {
    transition: all 1.5s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}

.slide-fade-enter,
.slide-fade-leave-to

/* .slide-fade-leave-active below version 2.1.8 */
    {
    transform: translateX(40px);
    opacity: 0;
}
</style>
