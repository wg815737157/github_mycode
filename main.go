package main

import (
	"algorithm/util"
	"errors"
	"fmt"
)

type Controller8090 struct {
}

//
//func (c *Controller8090) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//	a := `<!DOCTYPE html>
//<html>
//<head>
//    <meta charset="UTF-8">
//    <title>CORS跨域</title>
//    <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
//</head>
//<body>
//<div style="text-align:center;margin-top: 100px;font-size: 60px;color: brown;cursor: pointer;">
//    <span onclick="sendAjaxReq()">发送Ajax请求</span>
//</div>
//<script type="text/javascript">
//    function sendAjaxReq() {
//        $.ajax({
//            type: "GET",
//            // contentType: "application/json",
//            url: "http://localhost:8080/",
//			xhrFields: {
//                withCredentials: true
//            },
//            crossDomain: true,
//            success: function (message) {
//                console.log("成功！" + message);
//            },
//            error: function (a, b, c) {
//                console.log("失败！" + a.statusText);
//            }
//        });
//    }
//</script>
//</body>
//</html>`
//	res.Write([]byte(a))
//}

type Controller8080 struct {
}

//func (c *Controller8080) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//	fmt.Println(req.Header.Get("origin"))
//	res.Header().Set("Access-Control-Allow-Origin", req.Header.Get("origin"))
//	res.Header().Set("Access-Control-Allow-Credentials", "true")
//	_, err := res.Write([]byte(`localhost:8080`))
//	if err != nil {
//		fmt.Println(err)
//	}
//}

type Tmp struct {
}

//func (c *Tmp) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//	reqBodyByte, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBodyByte))
//	tmp, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("1" + string(tmp))
//	return
//}

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

// 二叉树节点数
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

// 中序遍历
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

// 删除链表节点
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

// 删除链表倒数第N个节点
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

// 删除链表倒数第N个节点
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

// 从下到上 需要改进同1没有算法进步
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

// 广度优先
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

// 深度优先
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

type Receiver struct {
	I int
}

func (r Receiver) setI(i int) {
	r.I = i
}

func main() {

	linkedHashMap := util.LinkedHashMapInit[string]()
	linkedHashMap.Add("zhangsan", "1")
	linkedHashMap.Add("lisi", "2")
	linkedHashMap.Add("wangwu", "3")
	linkedHashMap.Add("zhaoliu", "4")
	linkedHashMap.Print()
	return

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A", err, "B")
		}
	}()

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}

	array := []int{1, 2, 3, 4}
	for k, v := range array {
		fmt.Println(k, v)
	}

	mapA := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range array {
		fmt.Println(k, v)
	}
	for k, v := range mapA {
		fmt.Println(k, v)
	}

}
