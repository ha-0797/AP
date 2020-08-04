//ReadMe
//works
//when you give a the whole address of the directory and the json file, 
//when passing a list to search, pass each word as a separate command line arguement.


//does not work
//writing ..\mydirectory will not work
//writing [word1, word2, ..] in search will not work

//bugs
//could not read the words in a docx file 


const fs = require('fs')
const path = require('path')

var arr = [];
const index2 = (f, dest)=> {
	fs.readFile(f, (err, data) => {
		if(err){
			console.log("read")
		} else {
			data = data + ' '
			let lines = data.split("\r\n")
			let i = 1
			lines.forEach(line => {
				let words = line.split(" ")
				words.forEach(word => {
					word = word.replace(/[.,\/#!$%\^&\*;:{}=\-_`~()]/g,"");
					word = word.replace(/\s{2,}/g," ");
					word = word.toLowerCase()
					if(word.length > 3){
						let j = arr.filter(x => x.str == word)
						if(j != null){
							let index = arr.indexOf(j[0])
							if(index >= 0){
								let k = arr[index].list.filter(x => x.file == f)
								let index2 = arr[index].list.indexOf(k[0])
								if(index2 >= 0){
									if(arr[index].list[index2].line.indexOf(i) == -1){
										arr[index].list[index2].line.push(i)
									}
								} else {
									let obj = {
										file: f,
										line: [i]
									}
									arr[index].list.push(obj)
									let json = JSON.stringify(arr)
									fs.writeFile(dest, json, 'utf8',(data, err) => {
										if(err){
											console.log(`${err}`)
										} else {
										}
									})
								}
							} else {
								let obj2 = {
									file: f,
									line: [i]
								}
								let obj = {
									str: word,
									list: [obj2]
								}
								arr.push(obj)
								let json = JSON.stringify(arr)
								fs.writeFile(dest, json, 'utf8',(data, err) => {
									if(err){
										console.log(`${err}`)
									} else {
									}
								})
							}
						}
					}
				})
				++ i
			})
		}
	})
}
const index = (dest, srcFiles) => {
	fs.readdir(srcFiles, (err, files) => {
		if(err){
			console.log(`new file found, ${srcFiles}`)
			index2(srcFiles, dest)
		} else {
			console.log(files)
			files.forEach(f => {
				console.log(`new subfile found, ${f}`)
				index(dest, path.join(srcFiles, f))
			})
		}
	})
}

const print_json = obj => {
	obj.forEach(f =>{
		fs.readFile(f.file, (err, data) => {
			if(err){
				console.log("bad file or summit")
			} else {
				data = data + " "
				let lines = data.split("\r\n")
				f.line.forEach(i => {
					let j = 1
					lines.forEach(line => {
						if(i == j){
							console.log(line)
						}
						++ j
					})
				})	
			}
		})
	})
}


const search = (file, list_words) => {
	var data = JSON.parse(fs.readFileSync(file));
	list_words.forEach(word => {
		data.forEach(obj => {
			if(obj.str == word){
				print_json(obj.list)
			}
		})
	})
}

const run = () => {
	if(process.argv[2] == "index"){
		index(process.argv[3], process.argv[4])
	} else if(process.argv[2] == "search"){
		let arr1 = []
		process.argv.forEach(function (val, index, array) {
			if(index > 3){
				arr1.push(val)
			}
		})
		search(process.argv[3], arr1)
	}
}

run()