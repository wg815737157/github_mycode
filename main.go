package main

import (
	"bytes"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
)

type Controller8090 struct {
}

func (c *Controller8090) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	a := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>CORS跨域</title>
    <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
</head>
<body>
<div style="text-align:center;margin-top: 100px;font-size: 60px;color: brown;cursor: pointer;">
    <span onclick="sendAjaxReq()">发送Ajax请求</span>
</div>
<script type="text/javascript">
    function sendAjaxReq() {
        $.ajax({
            type: "GET",
            // contentType: "application/json",
            url: "http://localhost:8080/",
			xhrFields: {
                withCredentials: true
            },
            crossDomain: true,
            success: function (message) {
                console.log("成功！" + message);
            },
            error: function (a, b, c) {
                console.log("失败！" + a.statusText);
            }
        });
    }
</script>
</body>
</html>`
	res.Write([]byte(a))
}

type Controller8080 struct {
}

func (c *Controller8080) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header.Get("origin"))
	res.Header().Set("Access-Control-Allow-Origin", req.Header.Get("origin"))
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	_, err := res.Write([]byte(`localhost:8080`))
	if err != nil {
		fmt.Println(err)
	}
}

type Tmp struct {
}

func (c *Tmp) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	reqBodyByte, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBodyByte))
	tmp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("1" + string(tmp))
	return

	//time.Sleep(time.Second)
	//_, err := res.Write([]byte(`localhost:8080`))
	//if err != nil {
	//	fmt.Println(err)
	//}
}

type S struct {
	b float32
	A int32
	c float32
}

type BinaryTree struct {
	value  int
	slnode *BinaryTree
	srnode *BinaryTree
}

func invertBinaryTree(binaryTree *BinaryTree) {
	if binaryTree == nil {
		return
	}
	if binaryTree.slnode != nil || binaryTree.srnode != nil {
		if binaryTree.slnode != nil {
			fmt.Println("反转左侧")
			fmt.Println(binaryTree.slnode.value)
		}
		if binaryTree.srnode != nil {
			fmt.Println("反转右侧")
			fmt.Println(binaryTree.srnode.value)
		}
		tmpNode := binaryTree.slnode
		binaryTree.slnode = binaryTree.srnode
		binaryTree.srnode = tmpNode
		invertBinaryTree(binaryTree.slnode)
		invertBinaryTree(binaryTree.srnode)
	}

}

//二叉树节点数
func binaryTreeLength(binaryTree *BinaryTree) int {
	if binaryTree == nil {
		return 0
	}
	length := 1
	if binaryTree.slnode != nil {
		length += binaryTreeLength(binaryTree.slnode)
	}
	if binaryTree.srnode != nil {
		length += binaryTreeLength(binaryTree.srnode)
	}
	return length
}

//中序遍历
func printBinaryTree(binaryTree *BinaryTree) {
	if binaryTree == nil {
		return
	}
	printBinaryTree(binaryTree.slnode)

	fmt.Printf("%d \n", binaryTree.value)

	printBinaryTree(binaryTree.srnode)
}

func addNodeBinaryTree(binaryTree *BinaryTree, value int) error {
	if binaryTree == nil {
		return errors.New("tree is nil")
	}
	if value < binaryTree.value {
		if binaryTree.slnode == nil {
			binaryTreeNode := &BinaryTree{}
			binaryTreeNode.value = value
			binaryTree.slnode = binaryTreeNode
			return nil
		}
		addNodeBinaryTree(binaryTree.slnode, value)
		return nil
	}
	if binaryTree.srnode == nil {
		binaryTreeNode := &BinaryTree{}
		binaryTreeNode.value = value
		binaryTree.srnode = binaryTreeNode
		return nil
	}
	addNodeBinaryTree(binaryTree.srnode, value)
	//等值过滤
	return nil
}

//递增数组 求和
func twoSum(nums []int, target int) ([]interface{}, error) {
	if len(nums) < 2 {
		return nil, errors.New("length too short")
	}
	low := 0
	high := len(nums) - 1
	res := make([]interface{}, 0, 0)
	for low < high {
		if nums[low]+nums[high] == target {
			tmp := make([]int, 0, 0)
			//fmt.Println(nums[low], nums[high])
			tmp = append(tmp, low, high)
			res = append(res, tmp)
			low++
			high--
			continue
		}
		if nums[low]+nums[high] > target {
			high--
		} else {
			low++
		}
	}
	return res, nil
}

type LinkedListNode struct {
	preNode  *LinkedListNode
	value    interface{}
	nextNode *LinkedListNode
}

func (l *LinkedListNode) getLastNode() *LinkedListNode {
	headValue := l.value
	tmp := l
	for headValue != tmp.nextNode.value {
		tmp = l.nextNode
	}
	return tmp
}

func (l *LinkedListNode) addNode(value interface{}) {
	addNode := &LinkedListNode{nil, value, nil}
	if l == nil {
		l.preNode = addNode
		l.nextNode = addNode
		l.value = value
		return
	}
	lastNode := l.getLastNode()
	addNode.preNode = lastNode
	addNode.nextNode = l
	lastNode.nextNode = addNode
	l.preNode = addNode
	return
}
func (l *LinkedListNode) popNode() *LinkedListNode {
	if l == nil {
		return nil
	}
	tmp := l
	l = l.nextNode
	lastNode := l.getLastNode()
	lastNode.nextNode = l
	l.preNode = lastNode
	return tmp
}

func addList(head *LinkedListNode, node *LinkedListNode) (*LinkedListNode, error) {
	if head == nil {
		return node, nil
	}
	tmp := head
	for tmp.nextNode != nil {
		tmp = tmp.nextNode
	}
	tmp.nextNode = node
	return head, nil
}


//删除链表节点
func deleteN(head *LinkedListNode, n int) error {
	if head == nil {
		return errors.New("链表空")
	}
	if n == 0 {
		return errors.New("要删除的是链表的头节点")
	}
	tmp := head
	tmpNext := head.nextNode
	for i := 1; i < n; i++ {
		if tmp == nil {
			return errors.New("要删除的节点超过实际节点")
		}
		tmpNext = tmp
		tmp = tmp.nextNode
	}
	if tmp == nil {
		tmpNext.nextNode = nil
		return nil
	}
	tmp.nextNode = tmpNext.nextNode
	return nil
}

//删除链表倒数第N个节点
func deleteDescN_1(head *LinkedListNode, n int) error {
	if head == nil {
		return errors.New("链表空")
	}
	len := 1
	tmp := head
	for tmp.nextNode != nil {
		tmp = tmp.nextNode
		len++
	}
	if len < n {
		return errors.New("超过链表长度")
	}
	m := len - n
	err := deleteN(head, m)
	if err != nil {
		return err
	}
	return nil
}

//删除链表倒数第N个节点
func deleteDescN_2(head *LinkedListNode, n int) error {
	if head == nil || n == 0 {
		return errors.New("空链表或0")
	}
	i := 0
	tmp := head
	deleteNodePre := head
	deleteNode := head
	for tmp != nil {
		i++
		if i > n {
			deleteNodePre = deleteNode
			deleteNode = deleteNode.nextNode
		}
		tmp = tmp.nextNode
	}
	if i <= n {
		return errors.New("删除节点大于等于链表节点数")
	}
	deleteNodePre.nextNode = deleteNode.nextNode
	return nil
}

func printList(head *LinkedListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.value)
	printList(head.nextNode)

}

func getDept(binaryTree *BinaryTree) int {
	if binaryTree == nil {
		return 0
	}
	slNodeLen := getDept(binaryTree.slnode)
	srNodeLen := getDept(binaryTree.srnode)
	if slNodeLen > srNodeLen {
		return slNodeLen + 1
	}
	return srNodeLen + 1
}

// 自上向下
func isBalanceTree(binaryTree *BinaryTree) bool {
	if binaryTree == nil {
		return true
	}
	slNodeLen := getDept(binaryTree.slnode)
	srNodeLen := getDept(binaryTree.srnode)
	if slNodeLen-srNodeLen > 1 || slNodeLen-srNodeLen < -1 {
		return false
	}
	result := isBalanceTree(binaryTree.slnode)
	if result == false {
		return false
	}
	result = isBalanceTree(binaryTree.srnode)
	if result == false {
		return false
	}
	return true
}

//从下到上 需要改进同1没有算法进步
func isBalanceTree_2(binaryTree *BinaryTree) int {
	if binaryTree == nil {
		return 0
	}
	slNodeLen := isBalanceTree_2(binaryTree.slnode)
	srNodeLen := isBalanceTree_2(binaryTree.srnode)
	if slNodeLen < 0 {
		fmt.Println("返回左侧", binaryTree.slnode.value)
		return -1
	}
	if srNodeLen < 0 {
		fmt.Println("返回右侧", binaryTree.srnode.value)
		return -1
	}
	if slNodeLen-srNodeLen > 1 {
		fmt.Println("不平衡节点是", binaryTree.value)
		return -1
	}
	if srNodeLen-slNodeLen > 1 {
		fmt.Println("不平衡节点是", binaryTree.value)
	}
	if slNodeLen > srNodeLen {
		return slNodeLen + 1
	}
	return srNodeLen + 1
}

//广度优先
func BFS(binaryTreeSlice []*BinaryTree, routeArray []int, target int, resArray *[][]int) {
	if len(binaryTreeSlice) == 0 {
		return
	}
	tmp := make([]*BinaryTree, 0, 0)
	for _, binaryTree := range binaryTreeSlice {
		tmpValue := binaryTree.value
		routeArray = append(routeArray, tmpValue)
		if tmpValue == target {
			*resArray = append(*resArray, routeArray)
			continue
		}
		if binaryTree.slnode != nil {
			tmp = append(tmp, binaryTree.slnode)
		}
		if binaryTree.srnode != nil {
			tmp = append(tmp, binaryTree.srnode)
		}
	}
	BFS(tmp, routeArray, target, resArray)
	return
}

//深度优先
func DFS(binaryTree *BinaryTree, routeArray []int, target int, resArray *[][]int) {
	if binaryTree == nil {
		return
	}
	tmp := binaryTree.value
	routeArray = append(routeArray, tmp)
	if tmp == target {
		*resArray = append(*resArray, routeArray)
		return
	}
	if binaryTree.slnode != nil {
		DFS(binaryTree.slnode, routeArray, target, resArray)
	}
	if binaryTree.srnode != nil {
		DFS(binaryTree.srnode, routeArray, target, resArray)
	}
	return
}
func zero(ptr [32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	q := &Queue{}
	if root == nil {
		return nil
	}
	q.Add(root)
	for q.Len() != 0 {
		len := q.Len()
		list := make([]int, len)
		for i := 0; i < len; i++ {
			treeNode := q.Remove().(*TreeNode)
			list[i] = treeNode.Val
			if treeNode.Left != nil {
				q.Add(treeNode.Left)
			}
			if treeNode.Right != nil {
				q.Add(treeNode.Right)
			}
		}
		ret = append(ret, list)
	}
	return ret
}

type Receiver struct {
	I int
}

func (r Receiver) setI(i int) {
	r.I = i
}

func main() {
	s := [...]int{100: 1}
	fmt.Println(len(s), s)

	//node1 := &TreeNode{9, nil, nil}
	//node2 := &TreeNode{20, nil, nil}
	//root := &TreeNode{1, node1, node2}
	//fmt.Println(levelOrder(root))

	//cmdInterface := exec.Command("ls")
	//result, err := cmdInterface.CombinedOutput()
	//if err != nil {
	//	fmt.Println("ls err")
	//}
	//files := string(result)
	//fileNames := strings.Split(files, "\n")
	//for _, fileName := range fileNames {
	//	fmt.Println(fileName)
	//	fileInfo, err := os.Stat(fileName)
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//	if fileInfo.IsDir() {
	//
	//	}
	//}

	//c := make(chan int, 0)
	//
	//select {
	//case <-c:
	//	fmt.Println("<-chan")
	//default:
	//	fmt.Println("default")
	//}

	//binaryTree := &BinaryTree{value: 10}
	//addNodeBinaryTree(binaryTree, 2)
	//addNodeBinaryTree(binaryTree, 4)
	//addNodeBinaryTree(binaryTree, 555)
	//addNodeBinaryTree(binaryTree, 6)
	//ret := isBalanceTree_2(binaryTree)
	//if ret < 0 {
	//	fmt.Println("不平衡树")
	//	return
	//}
	//fmt.Println("平衡树")

	//target := 2
	//resArray := make([][]int, 0, 0)
	//routeArray := make([]int, 0, 0)
	//
	//tmp := make([]*BinaryTree, 0, 0)
	//tmp = append(tmp, binaryTree)
	//BFS(tmp, routeArray, target, &resArray)
	//fmt.Println(resArray)

	//
	//printBinaryTree(binaryTree)
	//
	//fmt.Println(isBalanceTree(binaryTree))

	//printBinaryTree(binaryTree)
	//length := binaryTreeLength(binaryTree)
	//fmt.Println(length)
	//fmt.Println("反转")
	//invertBinaryTree(binaryTree)

	//

	//
	//res := permute([]int{1, 2, 3, 4, 5})
	//
	//fmt.Println(res)
	//head := &ListNode{0, nil}
	//node1 := &ListNode{1, nil}
	//node2 := &ListNode{2, nil}
	//node3 := &ListNode{3, nil}
	//node4 := &ListNode{4, nil}
	//node5 := &ListNode{5, nil}
	//node6 := &ListNode{6, nil}
	//
	//head, err := addList(head, node1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//head, err = addList(head, node2)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//head, err = addList(head, node3)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//head, err = addList(head, node4)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//head, err = addList(head, node5)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//head, err = addList(head, node6)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//
	//
	//err = deleteDescN_2(head, 7)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//printList(head)

	//err = deleteN(head, 0)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//printList(head)

	//err = deleteDescN_1(head, 6)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//newHead := reverseList(head)
	//printList(newHead)

	//printList(revertHead)

	//tmpNums := make([]int, 0, 0)
	//tmpNums = append(tmpNums, 1, 4, 3, 5, 2, 7, 8, 9, 6)
	//
	//nums := quickSort(tmpNums)
	//fmt.Println(nums)
	////return
	//res, err := twoSum(nums, 5)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(res)

	//
	//svc := &http.Server{}
	//svc.Addr = ":8899"
	//svc.Handler = &Tmp{}
	////svc.WriteTimeout = time.Millisecond
	//err := svc.ListenAndServe()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//inputBytes, err := ioutil.ReadFile("1.txt")
	//
	//for _, v := range inputBytes {
	//	fmt.Printf("%x\n", v)
	//}
	//fmt.Printf("分隔符\n")
	//
	//str := string(inputBytes)
	//
	//runeStr := []rune(str)
	//
	//for _, v := range runeStr {
	//	fmt.Printf("%x\n", v)
	//}
	//return
	//
	//res, err := http.Get("https://account.jetbrains.com")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for k, v := range res.Header {
	//	fmt.Println(k)
	//	fmt.Println(v)
	//}
	//resByte, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(resByte))
	//return
	//go func() {
	//	c8080 := &Controller8080{}
	//	err := http.ListenAndServe(":8080", c8080)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	//
	//c8090 := &Controller8090{}
	//err = http.ListenAndServe(":8090", c8090)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//a := `12`
	//
	//var Aer interface{}
	//err := json.Unmarshal([]byte(a), &Aer)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//if Aer == nil {
	//	fmt.Println("Aer is null")
	//	return
	//}
	//
	//if mapA, ok := Aer.(map[string]interface{}); ok {
	//	for k, v := range mapA {
	//		fmt.Println(k, v)
	//	}
	//	return
	//}
	//if intA, ok := Aer.(float64); ok {
	//	fmt.Println(intA)
	//	return
	//}
	//if stringA, ok := Aer.(string); ok {
	//	fmt.Println(stringA)
	//	return
	//}

	//a := 1
	//
	//p := &a
	//
	//fmt.Printf("%p \n", &a)
	//
	//fmt.Printf("%p \n", p)
	//fmt.Printf("%p \n", &p)
	//fmt.Printf("%d \n", *p)

	//demo := &controller.Demo{}
	//http.HandleFunc("/", demo.Test)
	//err := http.ListenAndServe(":8090", nil)
	//if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//}
	//demo := &controller.Demo{}
	//err := http.ListenAndServe(":8090", demo)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
}
